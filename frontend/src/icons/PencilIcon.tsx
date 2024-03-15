import React from "react";

const PencilIcon = ({ stroke = "#9DA4AE", fill = "none", size = 20, ...props }: any) => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width={size} height={size} fill={fill}>
      <path
        stroke={stroke}
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.667}
        d="M17.5 17.5h-6.666m-8.75.417 4.624-1.779c.296-.114.444-.17.582-.245.123-.066.24-.142.35-.228.124-.096.236-.208.46-.432l9.4-9.4A2.357 2.357 0 1 0 14.167 2.5l-9.4 9.4c-.224.224-.336.336-.432.46a2.5 2.5 0 0 0-.228.35c-.074.138-.131.286-.245.582l-1.779 4.625Zm0 0 1.715-4.46c.122-.319.184-.478.289-.551a.417.417 0 0 1 .316-.067c.126.024.246.145.488.386l1.883 1.883c.242.242.362.362.386.488a.417.417 0 0 1-.067.316c-.073.105-.232.167-.551.29l-4.46 1.715Z"
      />
    </svg>
  );
}

export default PencilIcon;
