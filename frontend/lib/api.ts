import { Book, ScrapeRequest } from "@/types/types";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

class GoScraperAPI {
  private baseUrl: string;

  constructor(baseUrl: string = API_BASE_URL) {
    this.baseUrl = baseUrl;
  }

  async scrapeBook(request: ScrapeRequest): Promise<Book[]> {
    const response = await fetch(`${this.baseUrl}/scrape/books`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    });
    if (!response.ok) {
      throw new Error(`Failed to scrape books: ${response.statusText}`);
    }
    return response.json();
  }
  async getCategories(): Promise<string[]> {
    const response = await fetch(`${this.baseUrl}/books/categories`, {
      headers: { "Content-Type": "application/json" },
    });
    if (!response.ok) {
      throw new Error(`Failed to get categories: ${response.statusText}`);
    }
    return response.json();
  }
}

export const api = new GoScraperAPI();
