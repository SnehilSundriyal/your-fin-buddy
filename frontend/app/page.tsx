'use client';

import { useEffect, useState } from 'react';

export default function Home() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080')
        .then(res => res.json())
        .then(data => {
          setData(data);
          setLoading(false);
        })
        .catch(err => {
          setError(err.message);
          setLoading(false);
        });
  }, []);

  if (loading) return <div className="p-8">Loading...</div>;
  if (error) return <div className="p-8 text-red-500">Error: {error}</div>;

  return (
      <div className="p-8">
        <h1 className="text-2xl font-bold mb-4">Next.js + Go Backend</h1>
        <pre className="bg-gray-900 p-4 rounded">
        {JSON.stringify(data, null, 2)}
      </pre>
      </div>
  );
}