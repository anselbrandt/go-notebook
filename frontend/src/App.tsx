import { useEffect, useState } from "react";
import DraggableList from "./components/DraggableList";
import { Note } from "./types";

function App() {
  const [data, setData] = useState<Note[]>();

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:9090/notes");
      const json = await response.json();
      setData(json);
    };
    fetchData();
  }, []);
  return (
    <div className="flex fill center">
      {data && <DraggableList items={data} />}
    </div>
  );
}

export default App;
