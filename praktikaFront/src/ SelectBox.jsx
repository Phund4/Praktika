import React from 'react';
import './styles/select-box.css';

export function SelectBox({ filter, setFilter, standartValue, valueArr }) {
  return (
    <div className="select-box">
      <select
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
      >
        <option value="">{standartValue}</option>
        {valueArr.map((el, index) => <option key={`option-${index}`} value={el}>{el}</option>)}
      </select>
    </div>
  );
}