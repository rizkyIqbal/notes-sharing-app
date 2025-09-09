"use client";

import { usePathname } from "next/navigation";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/visitor/app-sidebar";

export default function LayoutWrapper({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();
  const isAuthPage = pathname === "/login" || pathname === "/register";

  return (
    <SidebarProvider>
      {!isAuthPage && <AppSidebar />}
      {!isAuthPage && <SidebarTrigger />}
      {children}
    </SidebarProvider>
  );
}
