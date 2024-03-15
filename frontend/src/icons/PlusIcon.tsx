import React from 'react';

const PlusIcon = ({ stroke = "#6C737F", fill = "none", size = 20, ...props }: any) => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width={size} height={size} fill={fill}>
      <path
        stroke={stroke}
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={2}
        d="M10 4.167v11.666M4.167 10h11.666"
      />
    </svg>
  );
}

export default PlusIcon;
