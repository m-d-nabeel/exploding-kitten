import { create } from "zustand";
import { persist, createJSONStorage, devtools } from "zustand/middleware";

type GameStore = {
  username: string;
  score: number;
  setScore: (newScore: number) => void;
  setUsername: (newUsername: string) => void;
};

export const useGameStore = create<GameStore>()(
  devtools(
    persist(
      (set) => ({
        username: "",
        score: 0,
        setScore: (newScore: number) => set({ score: newScore }),
        setUsername: (newUsername: string) => set({ username: newUsername }),
      }),
      {
        name: "exploding-kitten-storage",
        storage: createJSONStorage(() => sessionStorage),
      },
    ),
  ),
);
