import { Footer } from "@/components/Footer";
import { GameUI } from "@/components/GameUI";
import { Navbar } from "@/components/Header";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogTitle
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { useGameStore } from "@/hooks/use-store";
import { createFileRoute } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createFileRoute("/game/")({
  component: Game,
});

function Game() {
  const { username } = useGameStore();
  const [isModalOpen, setIsModalOpen] = useState(username.length === 0);
  const [inputValue, setInputValue] = useState("");
  const handleDialogToggle = () => {
    if (username.length === 0) {
      setIsModalOpen(true);
    } else {
      setIsModalOpen((prev) => !prev);
    }
  };

  const handleSubmit = () => {
    if (inputValue.length > 0) {
      useGameStore.setState({ username: inputValue });
      setIsModalOpen(false);
    }
  };

  return (
    <main className="flex min-h-screen w-full flex-col bg-gradient-to-br from-orange-400 via-orange-200 to-orange-400">
      <Navbar />
      <Dialog open={isModalOpen} onOpenChange={handleDialogToggle}>
        <DialogContent>
          <DialogTitle>Welcome!</DialogTitle>
          <Input
            type="text"
            min={3}
            max={64}
            name="username"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            autoFocus
            placeholder="Enter your unique username"
          />
          <Button className="ml-auto" onClick={handleSubmit}>
            OK
          </Button>
        </DialogContent>
      </Dialog>
      {!isModalOpen && (
        <div className="grid place-items-center">
          <GameUI />
        </div>
      )}
      <Footer />
    </main>
  );
}
