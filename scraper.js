"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const puppeteer_1 = __importDefault(require("puppeteer"));
const fs_1 = __importDefault(require("fs"));
function scrapeWebsite() {
    return __awaiter(this, void 0, void 0, function* () {
        const browser = yield puppeteer_1.default.launch({ headless: false, timeout: 120000 }); // Set headless to false for debugging
        const page = yield browser.newPage();
        yield page.goto("https://steamcommunity.com/profiles/76561198080078959/inventoryhistory"); // Change this to your target URL
        while (true) {
            yield page.waitForSelector("#inventory_history_table", { timeout: 120000 });
            const trades = yield page.evaluate(() => {
                const tradeRows = document.querySelectorAll(".tradehistoryrow");
                return Array.from(tradeRows).map((row) => {
                    var _a, _b, _c, _d, _e, _f;
                    // Get trade date & timestamp
                    const dateElement = row.querySelector(".tradehistory_date");
                    const date = ((_b = (_a = dateElement === null || dateElement === void 0 ? void 0 : dateElement.childNodes[0]) === null || _a === void 0 ? void 0 : _a.textContent) === null || _b === void 0 ? void 0 : _b.trim()) || "Unknown Date";
                    const time = ((_d = (_c = row.querySelector(".tradehistory_timestamp")) === null || _c === void 0 ? void 0 : _c.textContent) === null || _d === void 0 ? void 0 : _d.trim()) ||
                        "Unknown Time";
                    // Get trade description
                    const description = ((_f = (_e = row === null || row === void 0 ? void 0 : row.querySelector(".tradehistory_event_description")) === null || _e === void 0 ? void 0 : _e.textContent) === null || _f === void 0 ? void 0 : _f.trim()) || "Unknown Description";
                    // Get item data
                    const items = Array.from(row.querySelectorAll(".tradehistory_items_group")).map((itemGroup) => {
                        var _a, _b, _c, _d, _e, _f;
                        const plusMinus = ((_c = (_b = (_a = itemGroup
                            .closest(".tradehistory_items")) === null || _a === void 0 ? void 0 : _a.querySelector(".tradehistory_items_plusminus")) === null || _b === void 0 ? void 0 : _b.textContent) === null || _c === void 0 ? void 0 : _c.trim()) || "?";
                        const itemElement = itemGroup.querySelector("a.history_item, span.history_item");
                        const name = ((_e = (_d = itemElement === null || itemElement === void 0 ? void 0 : itemElement.querySelector(".history_item_name")) === null || _d === void 0 ? void 0 : _d.textContent) === null || _e === void 0 ? void 0 : _e.trim()) || "Unknown Item";
                        const image = ((_f = itemElement === null || itemElement === void 0 ? void 0 : itemElement.querySelector("img")) === null || _f === void 0 ? void 0 : _f.src) || "No Image";
                        return { plusMinus, name, image };
                    });
                    return { date, time, description, items };
                });
            });
            trades.forEach((trade) => {
                console.log(trade);
            });
            const unlockedContainers = trades.filter((trade) => trade.description.includes("Unlocked a container"));
            fs_1.default.writeFileSync("unlocked_container.json", JSON.stringify(unlockedContainers, null, 2), "utf-8");
            // Nächste Seite
            let loadMoreButton = yield page.$("div.load_more_history_area a");
            if (loadMoreButton) {
                yield waitForSecond();
                const initialRowCount = yield page.$$eval(".tradehistoryrow", (rows) => rows.length);
                yield loadMoreButton.click();
                yield page.waitForFunction((initialCount) => document.querySelectorAll(".tradehistoryrow").length > initialCount, { timeout: 15000 }, // Increase timeout if Steam is slow
                initialRowCount);
                console.log("new history loaded!");
            }
            else {
                console.log("load more button not found!");
                break;
            }
        }
        yield browser.close();
    });
}
function waitForSecond() {
    return new Promise((resolve) => setTimeout(resolve, 4000));
}
scrapeWebsite();
