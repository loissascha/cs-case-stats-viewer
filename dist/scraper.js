"use strict";
var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __getProtoOf = Object.getPrototypeOf;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toESM = (mod, isNodeMode, target) => (target = mod != null ? __create(__getProtoOf(mod)) : {}, __copyProps(
  // If the importer is in node compatibility mode or this is not an ESM
  // file that has been converted to a CommonJS file using a Babel-
  // compatible transform (i.e. "__esModule" has not been set), then set
  // "default" to the CommonJS "module.exports" for node compatibility.
  isNodeMode || !mod || !mod.__esModule ? __defProp(target, "default", { value: mod, enumerable: true }) : target,
  mod
));
var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};

// scraper.ts
var import_puppeteer = __toESM(require("puppeteer"));
var import_fs = __toESM(require("fs"));
function scrapeWebsite() {
  return __async(this, null, function* () {
    const browser = yield import_puppeteer.default.launch({ headless: false, timeout: 12e4 });
    const page = yield browser.newPage();
    yield page.goto(
      "https://steamcommunity.com/profiles/76561198080078959/inventoryhistory"
    );
    while (true) {
      yield page.waitForSelector("#inventory_history_table", { timeout: 12e4 });
      const trades = yield page.evaluate(() => {
        const tradeRows = document.querySelectorAll(".tradehistoryrow");
        return Array.from(tradeRows).map((row) => {
          var _a, _b, _c, _d, _e, _f;
          const dateElement = row.querySelector(".tradehistory_date");
          const date = ((_b = (_a = dateElement == null ? void 0 : dateElement.childNodes[0]) == null ? void 0 : _a.textContent) == null ? void 0 : _b.trim()) || "Unknown Date";
          const time = ((_d = (_c = row.querySelector(".tradehistory_timestamp")) == null ? void 0 : _c.textContent) == null ? void 0 : _d.trim()) || "Unknown Time";
          const description = ((_f = (_e = row == null ? void 0 : row.querySelector(".tradehistory_event_description")) == null ? void 0 : _e.textContent) == null ? void 0 : _f.trim()) || "Unknown Description";
          const items = Array.from(
            row.querySelectorAll(".tradehistory_items_group")
          ).map((itemGroup) => {
            var _a2, _b2, _c2, _d2, _e2, _f2;
            const plusMinus = ((_c2 = (_b2 = (_a2 = itemGroup.closest(".tradehistory_items")) == null ? void 0 : _a2.querySelector(".tradehistory_items_plusminus")) == null ? void 0 : _b2.textContent) == null ? void 0 : _c2.trim()) || "?";
            const itemElement = itemGroup.querySelector(
              "a.history_item, span.history_item"
            );
            const name = ((_e2 = (_d2 = itemElement == null ? void 0 : itemElement.querySelector(".history_item_name")) == null ? void 0 : _d2.textContent) == null ? void 0 : _e2.trim()) || "Unknown Item";
            const image = ((_f2 = itemElement == null ? void 0 : itemElement.querySelector("img")) == null ? void 0 : _f2.src) || "No Image";
            return { plusMinus, name, image };
          });
          return { date, time, description, items };
        });
      });
      trades.forEach((trade) => {
        console.log(trade);
      });
      const unlockedContainers = trades.filter(
        (trade) => trade.description.includes("Unlocked a container")
      );
      import_fs.default.writeFileSync(
        "unlocked_container.json",
        JSON.stringify(unlockedContainers, null, 2),
        "utf-8"
      );
      let loadMoreButton = yield page.$("div.load_more_history_area a");
      if (loadMoreButton) {
        yield waitForSecond();
        const initialRowCount = yield page.$$eval(
          ".tradehistoryrow",
          (rows) => rows.length
        );
        yield loadMoreButton.click();
        yield page.waitForFunction(
          (initialCount) => document.querySelectorAll(".tradehistoryrow").length > initialCount,
          { timeout: 15e3 },
          // Increase timeout if Steam is slow
          initialRowCount
        );
        console.log("new history loaded!");
      } else {
        console.log("load more button not found!");
        break;
      }
    }
    yield browser.close();
  });
}
function waitForSecond() {
  return new Promise((resolve) => setTimeout(resolve, 4e3));
}
scrapeWebsite();
