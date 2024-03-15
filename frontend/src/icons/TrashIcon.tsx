import React from "react";

const TrashIcon = ({ stroke = "#FFA8A2", strokeWidth = 2, fill = "none", size = 20, ...props }: any) => {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" width={size} height={size} style={{ cursor: "pointer" }} viewBox="0 0 20 20" fill={fill}>
      <path
        stroke={stroke}
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={strokeWidth}
        d="M8 2.5h5M3 5h15m-1.667 0-.584 8.766c-.088 1.315-.132 1.973-.416 2.472a2.5 2.5 0 0 1-1.082 1.012c-.516.25-1.175.25-2.493.25H9.242c-1.318 0-1.977 0-2.493-.25a2.5 2.5 0 0 1-1.082-1.012c-.284-.5-.328-1.157-.416-2.472L4.667 5m4.166 3.75v4.167m3.334-4.167v4.167"
      />
    </svg>
  );
}

export default TrashIcon;
