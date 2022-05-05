import DraggableList from "./components/DraggableList";
import { Note } from "./types";

const mockData: Note[] = [
  {
    ID: 1,
    Contents: "Read War and Peace.",
    CreatedAt: 1651747944,
    UpdatedAt: 1651747944,
  },
  {
    ID: 2,
    Contents: "Buy bananas.",
    CreatedAt: 1651747957,
    UpdatedAt: 1651747957,
  },
  {
    ID: 3,
    Contents: "Pick up dry cleaning.",
    CreatedAt: 1651747975,
    UpdatedAt: 1651747975,
  },
  {
    ID: 4,
    Contents: "Sell bitcoin.",
    CreatedAt: 1651747988,
    UpdatedAt: 1651747988,
  },
];

function App() {
  return (
    <div className="flex fill center">
      <DraggableList items={mockData.map((item) => item.Contents)} />
    </div>
  );
}

export default App;
