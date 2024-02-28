import { Outlet, createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout")({
  component: LayoutComponent,
});

function LayoutComponent() {
  return (
    <div>
      <div>MainLayout</div>
      <div>
        <Outlet />
      </div>
    </div>
  );
}
