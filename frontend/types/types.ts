export interface Book {
  title: string;
  price: string;
  rating: string;
  in_stock: boolean;
  url: string;
  image_url: string;
}

export interface ScrapeRequest {
  category: string;
  page: number;
}
