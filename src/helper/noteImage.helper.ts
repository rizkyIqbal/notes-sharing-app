import api from "@/lib/services/api";
import { NoteImage } from "@/types/noteImage";

export async function uploadToGmbr(file: File) {
  const formData = new FormData();
  formData.append("file", file);

  const res = await fetch("https://api.gmbr.web.id/upload", {
    method: "POST",
    body: formData,
  });

  if (!res.ok) throw new Error("Failed to upload to gmbr");
  return res.json();
}

export async function saveImagesToNote(noteId: string, imagePaths: string[]) {
  const res = await api.post(`/notes/${noteId}/images`, {
    image_path: imagePaths,
  });
  return res.data;
}

export async function getNoteImages(noteId: string): Promise<NoteImage[]> {
  try {
    const res = await api.get(`/notes/${noteId}/images`);
    return res.data.data as NoteImage[];
  } catch (error) {
    console.error("Failed to fetch note images:", error);
    throw error;
  }
}

export async function deleteNoteImage(
  imageId: string,
  onSuccess?: (res: { message: string }) => void,
  onError?: (err: unknown) => void
) {
  try {
    const res = await api.delete(`/notes/images/${imageId}`);
    const data = await res.data;
    if (onSuccess) onSuccess(data);
    return data;
  } catch (err) {
    if (onError) onError(err);
    throw err;
  }
}
