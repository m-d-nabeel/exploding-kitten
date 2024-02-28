import { useNavigate } from "@tanstack/react-router";
import { Button } from "./ui/button";

export const StartScreen = () => {
  const navigate = useNavigate();
  return (
    <div className="mx-auto my-auto w-3/4 border bg-gradient-to-tl p-12 contrast-150">
      <div className="flex flex-col items-center justify-center gap-y-10">
        <img src="/src/assets/exploding-kitten.png" alt="icon" width={400} height={400} />
        <Button
          variant={"default"}
          size={"lg"}
          onClick={() => navigate({ from: "/", to: "/game" })}
        >
          START GAME
        </Button>
      </div>
    </div>
  );
};
