export function saveToJSON<T>(data: T[], filename: string): void {
  if (!data || data.length === 0) return;

  const blob = new Blob([JSON.stringify(data, null, 2)], {
    type: "application/json",
  });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");

  link.href = url;
  link.download = filename.endsWith(".json") ? filename : `${filename}.json`;

  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
}

export function saveToCSV<T extends object>(data: T[], filename: string): void {
  if (!data || data.length === 0) return;

  const headers = Object.keys(data[0]) as (keyof T)[];

  const csvRows = [
    headers.join(","),
    ...data.map((row) =>
      headers
        .map((header) => {
          const value = row[header] ?? "";
          const escaped = ("" + value).replace(/"/g, '""');
          return `"${escaped}"`;
        })
        .join(",")
    ),
  ];

  const blob = new Blob([csvRows.join("\n")], {
    type: "text/csv;charset=utf-8;",
  });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = filename.endsWith(".csv") ? filename : `${filename}.csv`;
  link.click();
  URL.revokeObjectURL(url);
}
