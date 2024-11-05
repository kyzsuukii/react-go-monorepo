import { useState } from "react";

interface PingResponse {
  message: string;
}

export default function App() {
  const [ping, setPing] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handlePing = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const res = await fetch("/api/ping");

      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }

      const data = (await res.json()) as PingResponse;
      setPing(data.message);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to fetch");
      setPing(null);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <h1>React + Go Monorepo</h1>
      <button onClick={handlePing} disabled={isLoading}>
        {isLoading ? "Loading..." : "Ping"}
      </button>
      {error && <p>{error}</p>}
      {ping && <p>Result: {ping}</p>}
    </div>
  );
}
