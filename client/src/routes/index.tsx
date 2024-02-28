import { Footer } from "@/components/Footer";
import { Navbar } from "@/components/Header";
import { StartScreen } from "@/components/StartScreen";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  const navigate = useNavigate({ from: "/" });
  const [username, setUsername] = useState("");

  useEffect(() => {
    const storedUsername = localStorage.getItem("exploding-kittens-username");
    if (storedUsername) {
      setUsername(storedUsername);
    }
  }, []);
  if (username) {
    navigate({ to: "/game", replace: true });
  }
  return (
    <main className="flex min-h-screen w-full flex-col bg-gradient-to-br from-orange-400 via-orange-200 to-orange-400">
      <Navbar />
      <StartScreen />
      <Footer />
    </main>
  );
}
