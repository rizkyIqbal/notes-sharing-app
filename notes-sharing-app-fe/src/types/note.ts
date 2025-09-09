export interface Note {
  id: string;
  title: string;
  content: string;
  user_id?: number;
  username?: string;
  created_at: string;
  updated_at: string;
}

export type NotesResponse = {
  status: number;
  message: string;
  data?: {
    limit: number;
    notes: Note[];
    page: number;
    total: number;
  };
};

export type NoteResponse = {
  status: number;
  message: string;
  data?: Note;
};
