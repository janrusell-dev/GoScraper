import { api } from "@/lib/api";
import { Book } from "@/types/types";
import { useState } from "react";

export function useScraper() {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState(false);

  const handleScrape = async (category: string, page: number) => {
    setLoading(true);
    try {
      const data = await api.scrapeBook({ category, page });
      setBooks(data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return { books, loading, handleScrape };
}
