import puppeteer from "puppeteer";
import fs from "fs";

async function scrapeWebsite() {
    const browser = await puppeteer.launch({
        headless: false,
        timeout: 120000,
        executablePath: "/usr/bin/chromium",
        args: [
            "--no-sandbox",
            "--disable-setuid-sandbox",
            "--ozone-platform=wayland",
        ],
    }); // Set headless to false for debugging
    const page = await browser.newPage();
    await page.goto(
        "https://steamcommunity.com/profiles/76561198080078959/inventoryhistory",
    ); // Change this to your target URL

    while (true) {
        await page.waitForSelector("#inventory_history_table", { timeout: 120000 });

        const trades = await page.evaluate(() => {
            const tradeRows = document.querySelectorAll(".tradehistoryrow");
            return Array.from(tradeRows).map((row) => {
                // Get trade date & timestamp
                const dateElement = row.querySelector(".tradehistory_date");
                const date =
                    dateElement?.childNodes[0]?.textContent?.trim() || "Unknown Date";
                const time =
                    row.querySelector(".tradehistory_timestamp")?.textContent?.trim() ||
                    "Unknown Time";

                // Get trade description
                const description =
                    row
                        ?.querySelector(".tradehistory_event_description")
                        ?.textContent?.trim() || "Unknown Description";

                // Get item data
                const items = Array.from(
                    row.querySelectorAll(".tradehistory_items_group"),
                ).map((itemGroup) => {
                    const plusMinus =
                        itemGroup
                            .closest(".tradehistory_items")
                            ?.querySelector(".tradehistory_items_plusminus")
                            ?.textContent?.trim() || "?";
                    const itemElement = itemGroup.querySelector(
                        "a.history_item, span.history_item",
                    );
                    const name =
                        itemElement
                            ?.querySelector(".history_item_name")
                            ?.textContent?.trim() || "Unknown Item";
                    const image = itemElement?.querySelector("img")?.src || "No Image";

                    return { plusMinus, name, image };
                });

                return { date, time, description, items };
            });
        });
        trades.forEach((trade) => {
            console.log(trade);
        });
        const unlockedContainers = trades.filter((trade) =>
            trade.description.includes("Unlocked a container"),
        );
        fs.writeFileSync(
            "/app/output/unlocked_container.json",
            JSON.stringify(unlockedContainers, null, 2),
            "utf-8",
        );

        // Nächste Seite
        let loadMoreButton = await page.$("div.load_more_history_area a");
        if (loadMoreButton) {
            await waitForSecond();
            const initialRowCount = await page.$$eval(
                ".tradehistoryrow",
                (rows) => rows.length,
            );
            await loadMoreButton.click();
            await page.waitForFunction(
                (initialCount) =>
                    document.querySelectorAll(".tradehistoryrow").length > initialCount,
                { timeout: 15000 }, // Increase timeout if Steam is slow
                initialRowCount,
            );
            console.log("new history loaded!");
        } else {
            console.log("load more button not found!");
            break;
        }
    }

    await browser.close();
}

function waitForSecond() {
    return new Promise((resolve) => setTimeout(resolve, 4000));
}

scrapeWebsite();
