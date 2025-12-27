import { api } from "@/lib/api";
import { useEffect, useState } from "react";

export function useCategories() {
  const [categories, setCategories] = useState<string[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadCategories = async () => {
      try {
        const data = await api.getCategories();
        setCategories(data);
      } catch (err) {
        console.error("failed to load categories:", err);
      } finally {
        setLoading(false);
      }
    };
    loadCategories();
  }, []);
  return { categories, loading };
}
