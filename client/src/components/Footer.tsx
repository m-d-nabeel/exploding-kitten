import { Link } from "@tanstack/react-router";

export const Footer = () => {
  return (
    <footer className="mt-auto flex h-16 items-center justify-center">
      <div className="flex items-center justify-center gap-x-2">
        <p>&copy; 2024</p>
        <Link to="/" className="underline">
          Exploding Kittens
        </Link>
      </div>
    </footer>
  );
};
