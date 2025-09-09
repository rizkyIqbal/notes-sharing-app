"use client";

import { ChevronUp, Home, Inbox, User2 } from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useEffect, useState } from "react";
import { fetchUserByID } from "@/helper/user.helper";
import { User } from "@/types/user";
import Link from "next/link";
import { Button } from "../ui/button";
import { logoutUser } from "@/helper/auth.helper";
import { useRouter } from "next/navigation";
import { AxiosError } from "axios";

const items = [
  {
    title: "All Notes",
    url: "/",
    icon: Home,
  },
  {
    title: "My Notes",
    url: "/mynotes",
    icon: Inbox,
  },
];

export function AppSidebar() {
  const router = useRouter();
  const [user, setUser] = useState<User | null | undefined>(undefined);
  useEffect(() => {
    async function loadUser() {
      try {
        await fetchUserByID(setUser);
      } catch (error: unknown) {
        if (error instanceof AxiosError) {
          if (error.response?.status === 401) {
            setUser(null);
          }
        } else {
          console.error("Unexpected error:", error);
        }
      }
    }
    loadUser();
  }, []);

  async function handleLogout() {
    try {
      await logoutUser();
      setUser(null);
      router.push("/login");
    } catch (error) {
      console.error("Logout failed:", error);
    }
  }

  if (user === undefined) return;

  return (
    <Sidebar>
      <SidebarHeader className="p-4 ">
        <SidebarMenu>
          <SidebarMenuItem>
            <p>Notes Sharing App</p>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Application</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            {user ? (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <SidebarMenuButton>
                    <User2 /> {user.username}
                    <ChevronUp className="ml-auto" />
                  </SidebarMenuButton>
                </DropdownMenuTrigger>
                <DropdownMenuContent
                  side="top"
                  className="w-[--radix-popper-anchor-width]"
                >
                  <DropdownMenuItem onClick={handleLogout}>
                    <span>Sign out</span>
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            ) : (
              <Link href="/login">
                <Button variant="outline" className="w-full">
                  Login
                </Button>
              </Link>
            )}
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  );
}
