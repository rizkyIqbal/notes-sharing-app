"use client";

import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Textarea } from "@/components/ui/textarea";
import {
  getNoteImages,
  deleteNoteImage,
  uploadToGmbr,
  saveImagesToNote,
} from "@/helper/noteImage.helper";
import { fetchNoteByID, updateNoteHelper } from "@/helper/notes.helper";
import { fetchUserByID } from "@/helper/user.helper";
import { Note } from "@/types/note";
import { NoteImage } from "@/types/noteImage";
import { User } from "@/types/user";
import { Image } from "antd";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { toast } from "sonner";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";

export default function UpdateNote() {
  const { id } = useParams();
  const [note, setNote] = useState<Note | null>(null);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [user, setUser] = useState<User | null>(null);
  const [images, setImages] = useState<NoteImage[]>([]);
  // const [loadingImage, setLoadingImage] = useState(true);
  const [loadingSubmit, setLoadingSubmit] = useState(false);
  const [selectedImage, setSelectedImage] = useState<NoteImage | null>(null);
  const [files, setFiles] = useState<File[]>([]);
  const router = useRouter();

  useEffect(() => {
    async function loadNote() {
      if (!id) return;
      try {
        const data = await fetchNoteByID(String(id), setNote);
        if (data) {
          setTitle(data.title);
          setContent(data.content);
        }
      } catch (error) {
        console.error(error);
      }
    }

    loadNote();
  }, [id]);

  useEffect(() => {
    async function checkAccess() {
      try {
        const noteData = await fetchNoteByID(String(id), setNote);
        const userData = await fetchUserByID(setUser);

        if (noteData?.user_id !== userData?.id) {
          router.replace("/403"); // Not Authorized
        }
      } catch (err) {
        console.error(err);
        router.replace("/403");
      }
    }

    async function fetchImages() {
      try {
        const res = await getNoteImages(String(id));
        setImages(res);
      } catch (err) {
        console.error("Error fetching images:", err);
      }
      // finally {
      //   setLoadingImage(false);
      // }
    }

    fetchImages();
    checkAccess();
  }, [id, router]);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoadingSubmit(true);
    try {
      await updateNoteHelper(
        String(id),
        { title, content },
        async () => {
          try {
            if (files.length > 0) {
              const uploaded = await Promise.all(
                files.map((file) => uploadToGmbr(file))
              );

              const imagePaths = uploaded.map((res) => res.custom_url);

              await saveImagesToNote(String(id), imagePaths);

              console.log("Images saved:", imagePaths);
            }

            toast.success("Images uploaded!");
          } catch (err) {
            toast.error("Error uploading images");
            console.error(err);
          }

          toast.success("Note Updated!");
          setFiles([]);
          (document.getElementById("picture") as HTMLInputElement).value = "";
        },
        (error) => {
          toast.error("Error updating note: " + error.message);
        }
      );
    } finally {
      setLoadingSubmit(false);
    }
  }

  async function handleDeleteImage() {
    if (!selectedImage) return;
    try {
      await deleteNoteImage(selectedImage.id, () => {
        toast.success("Image deleted!");
        setImages((prev) => prev.filter((img) => img.id !== selectedImage.id));
        setSelectedImage(null);
      });
    } catch (err: unknown) {
      if (err instanceof Error) {
        toast.error("Error deleting image: " + err.message);
      } else {
        toast.error("Error deleting image");
      }
    }
  }

  if (!note) return <p>Loading...</p>;

  return (
    <div className="font-sans min-h-screen p-8 w-full">
      <p className="text-3xl">Note - Update Note</p>
      <Separator className="my-4" />
        <p>Author : {user?.username}</p>
      <div className="flex w-full gap-6 mt-6">
        <form onSubmit={handleSubmit} className="min-w-sm">
          <div className="mb-5">
            <label
              htmlFor="username"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Title
            </label>
            <input
              type="text"
              id="text"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Insert your title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </div>
          <div className="mb-5">
            <label
              htmlFor="Content"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Content
            </label>
            <Textarea
              value={content}
              onChange={(e) => setContent(e.target.value)}
            />
          </div>
          <div className="mb-5">
            <label
              htmlFor="picture"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Picture
            </label>
            <Input
              multiple
              id="picture"
              type="file"
              accept="image/*"
              onChange={(e) =>
                setFiles(e.target.files ? Array.from(e.target.files) : [])
              }
            />
          </div>
          <Button className="w-full mt-2" disabled={loadingSubmit}>
            {loadingSubmit ? "Submitting..." : "Submit"}
          </Button>
        </form>

        {/* Images Section */}
        <div>
          <p className="text-lg font-medium">Images</p>
          <p className="text-sm text-gray-500">*Click to delete images</p>
          <div className="flex flex-wrap gap-4 mt-4">
            {(images ?? []).map((img) => (
              <Dialog key={img.id}>
                <DialogTrigger asChild>
                  <Image
                    preview={false}
                    width={200}
                    src={img.image_path}
                    alt="note image"
                    className="cursor-pointer rounded-md"
                    onClick={() => setSelectedImage(img)}
                  />
                </DialogTrigger>
                <DialogContent>
                  <DialogHeader>
                    <DialogTitle>Delete Image</DialogTitle>
                    <DialogDescription>
                      Are you sure you want to delete this image?
                    </DialogDescription>
                  </DialogHeader>
                  <DialogFooter>
                    <DialogClose asChild>
                      <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button variant="destructive" onClick={handleDeleteImage}>
                      Delete
                    </Button>
                  </DialogFooter>
                </DialogContent>
              </Dialog>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
