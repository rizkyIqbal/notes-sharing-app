"use client";

import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { deleteNoteHelper, fetchNoteByID } from "@/helper/notes.helper";
import { fetchUserByID } from "@/helper/user.helper";
import { Note } from "@/types/note";
import { User } from "@/types/user";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { Image } from "antd";

import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { getNoteImages } from "@/helper/noteImage.helper";
import { NoteImage } from "@/types/noteImage";

export default function NoteDetail() {
  const { id } = useParams();
  const [note, setNote] = useState<Note | null>(null);
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState<User | null>(null);
  const [images, setImages] = useState<NoteImage[]>([]);
  const router = useRouter();

  useEffect(() => {
    async function loadNote() {
      if (!id) return;
      try {
        const data = await fetchNoteByID(String(id), setNote);
        console.log("Fetched note:", data);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    }

    async function getUserIDLogged() {
      try {
        const data = await fetchUserByID(setUser);
        console.log("Logged user:", data);
      } catch (error) {
        const err = error as { response?: { status: number }; message?: string };
        if (err.response?.status === 401) {
          console.log("No user logged in");
        } else {
          console.error("Failed to fetch logged user:", error);
        }
      }
    }

    async function fetchImages() {
      try {
        const res = await getNoteImages(String(id));
        setImages(res);
      } catch (err) {
        console.error("Error fetching images:", err);
      } finally {
        setLoading(false);
      }
    }

    fetchImages();
    loadNote();
    getUserIDLogged();
  }, [id]);

  async function deleteSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!note) return;

    await deleteNoteHelper(
      note.id, // pass the note id
      (deletedNote) => {
        console.log("Note deleted:", deletedNote);
        router.push("/note");
      },
      (error) => {
        alert("Error deleting note: " + error.message);
      }
    );
  }

  if (loading) return <p>Loading...</p>;
  if (!note) return <p>Note not found</p>;

  return (
    <div className="font-sans min-h-screen p-8 w-full">
      <div className="flex justify-between">
        <div>
          <p className="text-3xl">Note {note.title}</p>
        </div>
        {note.user_id === user?.id && (
          <div className="flex gap-4 mt-6">
            <Button
              onClick={() => router.push(`/note/edit/${note.id}`)}
              className="bg-blue-600 text-white"
            >
              Edit
            </Button>

            <Dialog>
              <form>
                <DialogTrigger asChild>
                  <Button className="bg-red-600 text-white">Delete</Button>
                </DialogTrigger>
                <DialogContent className="sm:max-w-[425px]">
                  <DialogHeader>
                    <DialogTitle>Delete Note</DialogTitle>
                    <DialogDescription>
                      Are you sure you wanna delete this Note?
                    </DialogDescription>
                  </DialogHeader>
                  <DialogFooter>
                    <DialogClose asChild>
                      <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button type="submit" onClick={deleteSubmit}>
                      Yes
                    </Button>
                  </DialogFooter>
                </DialogContent>
              </form>
            </Dialog>
          </div>
        )}
      </div>
      <Separator className="my-4" />

      <div className="w-full max-w-sm">
        <table className="w-full border-separate border-spacing-x-6 text-left">
          <tbody>
            <tr>
              <th className="font-medium text-gray-700">Created By</th>
              <td className="text-gray-900">{note.username}</td>
            </tr>
            <tr>
              <th className="font-medium text-gray-700">Last Updated</th>
              <td className="text-gray-900">
                {new Date(note.updated_at).toLocaleString()}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <Separator className="my-4" />
      <p>{note.content}</p>
      <div className="flex gap-4">
        {(images ?? []).map((img) => (
          <Image alt="image" width={200} key={img.id} src={img.image_path} />
        ))}
      </div>
    </div>
  );
}
