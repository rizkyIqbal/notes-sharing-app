import { NotesResponse, NoteResponse } from "@/types/note";
import api from "./api";

export const noteService = {
  async getNotes(page = 1, limit = 8): Promise<NotesResponse> {
    const response = await api.get<NotesResponse>("/notes", {
      params: { page, limit },
    });
    return response.data;
  },

  async getNoteByID(id: string): Promise<NoteResponse> {
    const response = await api.get<NoteResponse>(`/notes/${id}`);
    return response.data;
  },

  async getNoteByUserID(page = 1, limit = 8): Promise<NotesResponse> {
    const response = await api.get<NotesResponse>(`/notes/user`, {
      params: { page, limit },
      withCredentials: true,
    });
    return response.data;
  },

  async createNote(note: {
    title: string;
    content: string;
  }): Promise<NoteResponse> {
    const response = await api.post<NoteResponse>("/notes", note);
    return response.data;
  },

  async updateNote(
    id: string,
    note: { title: string; content: string }
  ): Promise<NoteResponse> {
    const response = await api.patch<NoteResponse>(`/notes/${id}`, note);
    return response.data;
  },

  async deleteNote(id: string) {
    const response = await api.delete<NoteResponse>(`/notes/${id}`, {
      withCredentials: true,
    });
    return response.data;
  },
};
