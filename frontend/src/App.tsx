import React, { useEffect, useState, useMemo } from "react";
import DraggableList from "./components/DraggableList";
import { Note } from "./types";
import { cssGrad } from "./utils/colors";
import styles from "./styles.module.css";

function App() {
  const [data, setData] = useState<Note[] | undefined>();
  const [value, setValue] = useState<string>();
  const [isShown, setIsShown] = useState(data === undefined);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:9090/notes");
      const json = await response.json();
      setData(json);
      if (json !== null) {
        setIsShown(false);
      }
    };
    fetchData();
  }, []);

  const createHandler = () => {
    setIsShown(true);
  };

  const aboutHandler = () => {
    alert("about");
  };

  const cancelHandler = () => {
    setIsShown(false);
    setValue(undefined);
  };

  const addHandler = () => {
    const addNote = async (note: {}) => {
      const response = await fetch("http://localhost:9090/notes", {
        method: "POST",
        body: JSON.stringify(note),
      });
      const newNote: Note = await response.json();
      if (data) {
        setData((prev) => [...(prev as Note[]), newNote]);
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

  const deleteHandler = (id: number) => {
    const deleteNote = async () => {
      await fetch(`http://localhost:9090/notes/${id}`, {
        method: "DELETE",
      });
    };
    deleteNote();
    setData((prev) => [...(prev as Note[])].filter((note) => note.ID !== id));
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
          onClick={createHandler}
          className={styles.circleButton}
          style={{
            background: `${useMemo(() => cssGrad(50, 70, 60), [])}`,
          }}
        >
          <div>╋</div>
        </div>
        <div
          onClick={aboutHandler}
          className={styles.circleButton}
          style={{
            background: `${useMemo(() => cssGrad(50, 70, 60), [])}`,
          }}
        >
          <div style={{ fontSize: "28px" }}>?</div>
        </div>
      </div>
      {isShown && (
        <div>
          <div className={styles.inputBox} style={{ display: "flex" }}>
            <textarea onChange={handleChange} />
          </div>
          <div style={{ display: "flex", justifyContent: "space-between" }}>
            <div onClick={cancelHandler} className={styles.actionButton}>
              <div>Cancel</div>
            </div>
            <div onClick={addHandler} className={styles.actionButton}>
              <div>Add</div>
            </div>
          </div>
        </div>
      )}
      {data && <DraggableList data={data} deleteHandler={deleteHandler} />}
    </div>
  );
}

export default App;
