"use client";

import { useCategories } from "@/hooks/useCategory";
import { useScraper } from "@/hooks/useScraper";
import { saveToCSV, saveToJSON } from "@/lib/utils/download";
import { useState } from "react";

export default function Home() {
  const { categories } = useCategories();
  const { books, loading, handleScrape } = useScraper();

  const [selectedCategory, setSelectedCategory] = useState("");
  const [page, setPage] = useState(1);
  const handleDownload = () => {
    const filename = selectedCategory
      ? `scrape-${selectedCategory}`
      : "scrape-all";
    saveToJSON(books, filename);
  };

  return (
    <div className="p-6 font-mono max-w-4xl mx-auto">
      <div className="border-2 border-black p-4 bg-gray-50 shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]">
        <h1 className="text-xl font-bold mb-4 uppercase tracking-tighter">
          GoScraper CLI Interface
        </h1>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
          <div className="flex flex-col">
            <label className="text-xs font-bold mb-1">
              CATEGORY (OPTIONAL)
            </label>
            <select
              className="border-2 border-black p-2 bg-white"
              value={selectedCategory}
              onChange={(e) => setSelectedCategory(e.target.value)}
            >
              <option value="">-- ALL CATEGORIES --</option>
              {categories.map((c) => (
                <option key={c} value={c}>
                  {c}
                </option>
              ))}
            </select>
          </div>

          <div className="flex flex-col">
            <label className="text-xs font-bold mb-1">PAGE</label>
            <input
              type="number"
              className="border-2 border-black p-2"
              value={page}
              onChange={(e) => setPage(Number(e.target.value))}
              min={1}
            />
          </div>
          <div className="flex items-end">
            <button
              onClick={() => handleScrape(selectedCategory, page)}
              disabled={loading}
              className="w-full bg-black text-white font-bold p-2 hover:bg-gray-800 disabled:bg-gray-400"
            >
              {loading ? "EXECUTING..." : "RUN SCRAPER"}
            </button>
          </div>
        </div>

        {books.length > 0 && (
          <div className="mb-4 flex justify-between items-center bg-yellow-200 p-2 border-2 border-black">
            <span className="text-sm font-bold">
              {books.length} BOOKS FOUND
            </span>
            <div className="flex gap-2">
              <button
                onClick={() => saveToJSON(books, `export-${Date.now()}`)}
                className="bg-white border-2 border-black px-3 py-1 text-xs font-bold hover:bg-gray-100"
              >
                GET JSON
              </button>
              <button
                onClick={() => saveToCSV(books, `export-${Date.now()}`)}
                className="bg-white border-2 border-black px-3 py-1 text-xs font-bold hover:bg-gray-100"
              >
                GET CSV
              </button>
            </div>
          </div>
        )}

        {/* Raw Data Preview */}
        <div className="bg-white border-2 border-black p-4 h-96 overflow-auto text-xs">
          {loading ? (
            <div className="animate-pulse">_ Fetching from Go server...</div>
          ) : books.length > 0 ? (
            <pre>{JSON.stringify(books, null, 2)}</pre>
          ) : (
            <span className="text-gray-400">
              // No data loaded. Run scraper to view results.
            </span>
          )}
        </div>
      </div>
    </div>
  );
}
