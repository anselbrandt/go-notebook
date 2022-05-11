import React, { useEffect, useState, useMemo, useRef } from "react";
import DraggableList from "./components/DraggableList";
import { Note } from "./types";
import { cssGrad } from "./utils/colors";
import styles from "./styles.module.css";

function App() {
  const textareaRef = useRef<HTMLDivElement>(null);
  const [data, setData] = useState<Note[] | undefined>();
  const [value, setValue] = useState<string>();
  const [isShown, setIsShown] = useState(data === undefined);
  const [order, setOrder] = useState<number[] | undefined>();

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:8080/notes");
      const json = await response.json();
      setData(json);
      if (json !== null) {
        setIsShown(false);
      }
    };
    fetchData();
  }, []);

  const handleCreate = () => {
    textareaRef.current?.scrollIntoView({ behavior: "smooth" });
    setIsShown(true);
  };

  const handleAbout = () => {
    alert("about");
  };

  const handleCancel = () => {
    setIsShown(false);
    setValue(undefined);
  };

  const handleAdd = () => {
    const addNote = async (note: {}) => {
      const response = await fetch("http://localhost:8080/notes", {
        method: "POST",
        body: JSON.stringify(note),
      });
      const newNote: Note = await response.json();
      if (data) {
        if (order) {
          setData((prev) => {
            return [...order.map((i) => [...(prev as Note[])][i]), newNote];
          });
        } else {
          setData((prev) => [...(prev as Note[]), newNote]);
        }
      } else {
        setData([newNote]);
      }
    };
    if (value !== undefined) {
      const note = { contents: value };
      addNote(note);
      setValue(undefined);
      setIsShown(false);
    }
  };

  const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    const input = event.target.value;
    setValue(input);
  };

  const handleDelete = (id: number) => {
    const deleteNote = async () => {
      await fetch(`http://localhost:8080/notes/${id}`, {
        method: "DELETE",
      });
    };
    deleteNote();
    setData((prev) => [...(prev as Note[])].filter((note) => note.ID !== id));
  };

  const handleSort = (newOrder: number[]) => {
    // api call to backend to update sort order
    setOrder(newOrder);
  };

  return (
    <div
      style={{
        marginTop: "50px",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <div
        style={{
          width: "320px",
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <div
          onClick={handleCreate}
          className={styles.circleButton}
          style={{
            background: `${useMemo(() => cssGrad(50, 70, 60), [])}`,
          }}
        >
          <div>╋</div>
        </div>
        <div
          onClick={handleAbout}
          className={styles.circleButton}
          style={{
            background: `${useMemo(() => cssGrad(50, 70, 60), [])}`,
          }}
        >
          <div style={{ fontSize: "28px" }}>?</div>
        </div>
      </div>

      {data && (
        <DraggableList
          items={data}
          handleDelete={handleDelete}
          handleSort={handleSort}
        />
      )}
      {isShown && (
        <div ref={textareaRef}>
          <div className={styles.inputBox} style={{ display: "flex" }}>
            <textarea onChange={handleChange} />
          </div>
          <div style={{ display: "flex", justifyContent: "space-between" }}>
            <div onClick={handleCancel} className={styles.actionButton}>
              <div>Cancel</div>
            </div>
            <div onClick={handleAdd} className={styles.actionButton}>
              <div>Add</div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
