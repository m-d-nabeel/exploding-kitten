export const Navbar = () => {
  return (
    <nav className="mb-auto flex h-16 w-full flex-col px-6 sm:px-12 md:px-16 lg:px-24">
      <div className="flex items-center justify-between">
        <span className="flex items-center justify-center gap-x-3">
          <div className="text-2xl">😼</div>
          <div className="text-base font-semibold">EXPLODING KITTENS</div>
        </span>
        <div className="">SOME THINGS</div>
      </div>
      <hr className="border-orange-950" />
    </nav>
  );
};
