import { noteService } from "@/lib/services/note.service";
import type { Note } from "@/types/note";

export async function fetchNotes(
  setNotes: (notes: Note[]) => void,
  page = 1,
  limit = 8
): Promise<{
  notes: Note[];
  total: number;
  page: number;
  limit: number;
} | null> {
  try {
    const data = await noteService.getNotes(page, limit);

    const result = {
      notes: data.data?.notes || [],
      total: data.data?.total || 0,
      page: data.data?.page || page,
      limit: data.data?.limit || limit,
    };

    setNotes(result.notes);
    return result;
  } catch (error) {
    console.error("Failed to fetch notes:", error);
    return null;
  }
}

export async function fetchNoteByID(
  id: string,
  setNote: (note: Note | null) => void
): Promise<Note | null> {
  try {
    const response = await noteService.getNoteByID(id);
    const note = response.data ?? null;
    setNote(note);
    return note;
  } catch (error) {
    console.error("Failed to fetch note by ID:", error);
    setNote(null);
    return null;
  }
}

export async function fetchNotesByUserID(
  setNotes: (notes: Note[]) => void,
  page = 1,
  limit = 8
): Promise<{
  notes: Note[];
  total: number;
  page: number;
  limit: number;
} | null> {
  try {
    const response = await noteService.getNoteByUserID(page, limit);
    const {
      notes,
      total,
      page: currentPage,
      limit: currentLimit,
    } = response.data ?? {
      notes: [],
      total: 0,
      page,
      limit,
    };

    setNotes(notes);
    return { notes, total, page: currentPage, limit: currentLimit };
  } catch (error) {
    console.error("Failed to fetch notes by user ID:", error);
    setNotes([]);
    return null;
  }
}

export async function createNoteHelper(
  newNote: { title: string; content: string },
  onSuccess?: (note: Note) => void,
  onError?: (error: any) => void
): Promise<Note | null> {
  try {
    const response = await noteService.createNote(newNote);
    if (response.data) {
      onSuccess?.(response.data);
      return response.data;
    }
    return null;
  } catch (error) {
    console.error("Failed to create note:", error);
    onError?.(error);
    return null;
  }
}

export async function updateNoteHelper(
  id: string,
  newNote: { title: string; content: string },
  onSuccess?: () => void,
  onError?: (error: any) => void
): Promise<Note | null> {
  try {
    const response = await noteService.updateNote(id, newNote);
    if (response.status === 200) {
      onSuccess?.();
      // return response.data;  
    }
    return null;
  } catch (error) {
    console.error("Failed to update note:", error);
    onError?.(error);
    return null;
  }
}

export async function deleteNoteHelper(
  id: string,
  onSuccess?: (note: Note) => void,
  onError?: (error: any) => void
): Promise<Note | null> {
  try {
    const response = await noteService.deleteNote(id);
    if (response.data) {
      onSuccess?.(response.data);
      return response.data;
    }
    return null;
  } catch (error) {
    console.error("Failed to delete note:", error);
    onError?.(error);
    return null;
  }
}
