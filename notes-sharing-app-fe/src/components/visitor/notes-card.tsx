"use client";

import { Separator } from "@/components/ui/separator";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { IoAttachSharp } from "react-icons/io5";
import { LuCalendarDays } from "react-icons/lu";
import type { Note } from "@/types/note";
import { useEffect, useState } from "react";
import { getNoteImages } from "@/helper/noteImage.helper";

type NotesCardProps = {
  note: Note;
};

export function NotesCard({ note }: NotesCardProps) {
  // const [loading, setLoading] = useState(true);
  const [totalImages, setTotalImages] = useState<number>(0);
  useEffect(() => {
    async function fetchImages() {
      try {
        const res = await getNoteImages(String(note.id));
        setTotalImages(res?.length ?? 0);
      } catch (err) {
        console.error("Error fetching images:", err);
      } 
      // finally {
      //   setLoading(false);
      // }
    }

    fetchImages();
  });

  return (
    <div className="bg-gray-100 w-full py-6 px-6 rounded-xl">
      <p className="font-bold">{note.title}</p>
      <p className="mt-2 line-clamp-3">{note.content}</p>
      <Separator className="my-4" />
      <div className="flex justify-between">
        <div className="flex gap-1">
          <div className="*:data-[slot=avatar]:ring-background flex -space-x-2 *:data-[slot=avatar]:ring-2 *:data-[slot=avatar]:grayscale">
            <Avatar>
              <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
            <Avatar>
              <AvatarImage src="https://github.com/leerob.png" alt="@leerob" />
              <AvatarFallback>LR</AvatarFallback>
            </Avatar>
            <Avatar>
              <AvatarImage
                src="https://github.com/evilrabbit.png"
                alt="@evilrabbit"
              />
              <AvatarFallback>ER</AvatarFallback>
            </Avatar>
          </div>
          <div className="flex items-center">
            <IoAttachSharp size={24} />
            <p>{totalImages}</p>
          </div>
        </div>
        <div className="flex items-center justify-center gap-1">
          <LuCalendarDays size={20} />
          <p className="text-xs">
            {note.created_at ? new Date(note.created_at).toLocaleString() : "—"}
          </p>
        </div>
      </div>
    </div>
  );
}
