import React, { useRef } from "react";
import { useSprings, animated, config } from "@react-spring/web";
import { useDrag } from "@use-gesture/react";
import styles from "../styles.module.css";
import { clamp, swap } from "../utils";
import { cssGrad } from "../utils/colors";
import Card from "./Card";
import { Note } from "../types";

const fn =
  (order: number[], active = false, originalIndex = 0, curIndex = 0, y = 0) =>
  (index: number) =>
    active && index === originalIndex
      ? {
          y: curIndex * 100 + y,
          scale: 1.1,
          zIndex: 1,
          shadow: 15,
          immediate: (key: string) => key === "zIndex",
          config: (key: string) =>
            key === "y" ? config.stiff : config.default,
        }
      : {
          y: order.indexOf(index) * 100,
          scale: 1,
          zIndex: 0,
          shadow: 1,
          immediate: false,
        };

export default function DraggableList({ items }: { items: Note[] }) {
  const order = useRef(items.map((_, index) => index)); // Store indicies as a local ref, this represents the item order
  const [springs, api] = useSprings(items.length, fn(order.current)); // Create springs, each corresponds to an item, controlling its transform, scale, etc.
  const bind = useDrag(({ args: [originalIndex], active, movement: [, y] }) => {
    const curIndex = order.current.indexOf(originalIndex);
    const curRow = clamp(
      Math.round((curIndex * 100 + y) / 100),
      0,
      items.length - 1
    );
    const newOrder = swap(order.current, curIndex, curRow);
    api.start(fn(newOrder, active, originalIndex, curIndex, y)); // Feed springs new style data, they'll animate the view without causing a single render
    if (!active) order.current = newOrder;
  });

  return (
    <div>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <div
          className={styles.circleButton}
          style={{
            background: `${cssGrad(50, 70, 60)}`,
          }}
        >
          <div>╋</div>
        </div>
        <div
          className={styles.circleButton}
          style={{
            background: `${cssGrad(50, 70, 60)}`,
          }}
        >
          <div style={{ fontSize: "28px" }}>?</div>
        </div>
      </div>
      <div>
        <div className={styles.inputBox}>Text Input</div>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <div className={styles.actionButton}>
            <div>Cancel</div>
          </div>
          <div className={styles.actionButton}>
            <div>Add</div>
          </div>
        </div>
      </div>
      <div className={styles.content} style={{ height: items.length * 100 }}>
        {springs.map(({ zIndex, shadow, y, scale }, i) => {
          return (
            <animated.div
              {...bind(i)}
              key={i}
              style={{
                background: `${cssGrad(50, 70, 60)}`,
                zIndex,
                boxShadow: shadow.to(
                  (s) => `rgba(0, 0, 0, 0.15) 0px ${s}px ${2 * s}px 0px`
                ),
                y,
                scale,
              }}
            >
              <Card item={items[i]} />
            </animated.div>
          );
        })}
      </div>
    </div>
  );
}
