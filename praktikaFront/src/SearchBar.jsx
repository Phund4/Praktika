import React from 'react';
import './styles/search-bar.css';

export function SearchBar({ searchTerm, setSearchTerm }) {
  return (
    <div className="search-bar">
      <input
        type="text"
        placeholder="Поиск вакансий..."
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
    </div>
  );
}