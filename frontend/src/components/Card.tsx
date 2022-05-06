import React from "react";
import { Note } from "../types";
import styles from "../styles.module.css";

interface CardProps {
  item: Note;
  deleteHandler: () => void;
}

const Card: React.FC<CardProps> = ({ item, deleteHandler }: CardProps) => {
  return (
    <div
      className={styles.card}
      style={{ display: "flex", justifyContent: "space-between" }}
    >
      <div>{item["Contents"]}</div>
      <div
        onClick={deleteHandler}
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
};

export default Card;
