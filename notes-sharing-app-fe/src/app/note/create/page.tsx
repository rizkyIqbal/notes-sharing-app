"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { Textarea } from "@/components/ui/textarea";
import { createNoteHelper } from "@/helper/notes.helper";
import { uploadToGmbr, saveImagesToNote } from "@/helper/noteImage.helper"; // batch save
import { useState } from "react";
import { toast } from "sonner";

export default function CreateNote() {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [files, setFiles] = useState<File[]>([]);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);

    try {
      await createNoteHelper(
        { title, content },
        async (note) => {
          if (files.length > 0) {
            const uploaded = await Promise.all(
              files.map((file) => uploadToGmbr(file))
            );
            const imagePaths = uploaded.map((res) => res.custom_url);
            await saveImagesToNote(note.id, imagePaths);
          }

          toast.success("Note Created!");

          setTitle("");
          setContent("");
          setFiles([]);
          (document.getElementById("picture") as HTMLInputElement).value = "";
        },
        (error) => {
          alert("Error creating note: " + error.message);
        }
      );
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="font-sans min-h-screen p-8 w-full">
      <p className="text-3xl">Note - Create Note</p>
      <Separator className="my-4" />
      <form onSubmit={handleSubmit} className="max-w-sm">
        <div className="mb-5">
          <label
            htmlFor="title"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Title
          </label>
          <input
            type="text"
            id="title"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
            placeholder="Insert your title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>

        <div className="mb-5">
          <label
            htmlFor="content"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Content
          </label>
          <Textarea
            id="content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
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

        <Button className="w-full mt-2" type="submit" disabled={loading}>
          {loading ? "Submitting..." : "Submit"}
        </Button>
      </form>
    </div>
  );
}
