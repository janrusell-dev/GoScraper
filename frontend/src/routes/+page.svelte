<script lang="ts">
  import { writable } from "svelte/store";

  let categories = ["All Books", "Travel", "Mystery", "Historical Fiction"];
  let selectedCategory = "All Books";
  let page = 1;
  let format = "json"; // json or csv
  let loading = writable(false);
  let books = writable<any[]>([]);

  async function scrapeBooks() {
    loading.set(true);
    books.set([]);

    const payload = {
      page: Number(page),
      category: selectedCategory === "All Books" ? "" : selectedCategory,
      format
    };

    try {
      const res = await fetch("http://localhost:8080/scrape/books", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });

      if (!res.ok) throw new Error("Failed to scrape books");

      if (format === "json") {
        const data = await res.json();
        books.set(data);
      } else if (format === "csv") {
        const blob = await res.blob();
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = "books.csv";
        a.click();
        window.URL.revokeObjectURL(url);
      }
    } catch (err) {
      console.error(err);
      alert("Error scraping books. See console.");
    } finally {
      loading.set(false);
    }
  }
</script>

<div class="p-8">
  <h1 class="text-3xl mb-6 font-bold">Scraper</h1>

  <div class="mb-4 space-x-4">
    <select bind:value={selectedCategory} class="border px-2 py-1">
      {#each categories as cat}
        <option value={cat}>{cat}</option>
      {/each}
    </select>

    <input type="number" min="1" bind:value={page} class="border px-2 py-1 w-20" placeholder="Page" />

    <select bind:value={format} class="border px-2 py-1">
      <option value="json">JSON</option>
      <option value="csv">CSV</option>
    </select>

    <button on:click={scrapeBooks} class="bg-blue-500 text-white px-4 py-1 rounded">
      Scrape
    </button>
  </div>

  {#if $loading}
    <p class="text-gray-500">Scraping books, please wait...</p>
  {/if}

  {#if $books.length > 0 && format === "json"}
    <table class="border-collapse border border-gray-300 w-full mt-4">
      <thead>
        <tr>
          <th class="border px-2 py-1">Title</th>
          <th class="border px-2 py-1">Price</th>
          <th class="border px-2 py-1">Rating</th>
          <th class="border px-2 py-1">URL</th>
          <th class="border px-2 py-1">ImageURL</th>
          <th class="border px-2 py-1">In Stock</th>
        </tr>
      </thead>
      <tbody>
        {#each $books as book}
          <tr>
            <td class="border px-2 py-1">{book.title}</td>
            <td class="border px-2 py-1">{book.price}</td>
            <td class="border px-2 py-1">{book.ratings}</td>
             <td class="border px-2 py-1 underline"><a href="{book.url}" target="_blank">
              {book.url}
            </a>
          </td>
            <td class="border px-2 py-1 underline"><img src="{book.image_url}" alt="{book.title}"
              loading="lazy" decoding="async" class="object-cover h-full w-full"/></td>
            <td class="border px-2 py-1">{book.in_stock ? "In Stock" : " Out of Stock"}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</div>
