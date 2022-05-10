import React from "react";
import { Note } from "../types";
import styles from "../styles.module.css";

interface CardProps {
  item: Note;
  handleDelete: (id: number) => void;
}

const Card: React.FC<CardProps> = ({ item, handleDelete }: CardProps) => {
  return (
    <div className={styles.card}>
      <div style={{ display: "flex", justifyContent: "space-between" }}>
        <div className={styles.cardMain}>{item["Contents"]}</div>

        <div
          onClick={() => handleDelete(item.ID)}
          className={styles.cardClose}
          style={{
            marginRight: "10px",
            fontSize: "24px",
          }}
        >
          ✕
        </div>
      </div>
      <div className={styles.cardSub}>{`${new Date(
        item.UpdatedAt * 1000
      ).toLocaleString()}`}</div>
    </div>
  );
};

export default Card;
