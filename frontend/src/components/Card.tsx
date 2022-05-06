import React from "react";
import { Note } from "../types";

export default function Card({ item }: { item: Note }) {
  return (
    <div>
      <div>{item["Contents"]}</div>
    </div>
  );
}
