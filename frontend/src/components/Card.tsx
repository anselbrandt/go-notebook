import React from "react";
import { Note } from "../types";
import styles from "../styles.module.css";

export default function Card({ item }: { item: Note }) {
  return (
    <div
      className={styles.card}
      style={{ display: "flex", justifyContent: "space-between" }}
    >
      <div>{item["Contents"]}</div>
      <div
        className={styles.cardClose}
        style={{
          marginRight: "10px",
          fontSize: "24px",
        }}
      >
        ✕
      </div>
    </div>
  );
}
