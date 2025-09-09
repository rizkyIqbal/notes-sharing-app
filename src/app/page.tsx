"use client";

import { Input } from "@/components/ui/input";
import { NotesCard } from "@/components/visitor/notes-card";
import { useEffect, useState } from "react";
import { Note } from "@/types/note";
import { fetchNotes } from "@/helper/notes.helper";

import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import Link from "next/link";

export default function Home() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [page, setPage] = useState(1);
  const [limit] = useState(8);
  const [totalPages, setTotalPages] = useState<number>(1);

  useEffect(() => {
    async function loadNotes() {
      try {
        const data = await fetchNotes(setNotes, page, limit);
        if (data?.total) {
          setTotalPages(Math.ceil(data.total / limit));
        }
      } catch (error) {
        console.error(error);
      }
    }
    loadNotes();
  }, [page, limit]);

  const handlePageClick = (newPage: number) => {
    setPage(newPage);
  };

  return (
    // <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
    <div className="font-sans min-h-screen p-8 w-full">
      {/* <main className="flex flex-col gap-[32px] row-start-2 items-center sm:items-start"> */}
      <main className="">
        <p className="text-3xl">Home</p>
        <div className="mt-4">
          <Input
            type="text"
            id="search"
            placeholder="Search Note"
            className="w-64"
          />
        </div>
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 mt-10 gap-4 min-w-full">
          {notes && notes.length > 0 ? (
            notes.map((note) => (
              <Link href={`/note/${note.id}`} key={note.id}>
                <NotesCard key={note.id} note={note} />
              </Link>
            ))
          ) : (
            <p className="col-span-full text-center text-gray-500 text-lg">
              No notes found
            </p>
          )}
        </div>
        {notes && notes.length > 0 && (
          <div className="mt-12 flex justify-center">
            <Pagination>
              <PaginationContent>
                <PaginationItem>
                  <PaginationPrevious
                    href="#"
                    onClick={(e) => {
                      e.preventDefault();
                      if (page > 1) setPage(page - 1);
                    }}
                  />
                </PaginationItem>

                {/* Generate page links dynamically */}
                {Array.from({ length: totalPages }, (_, i) => i + 1).map(
                  (p) => (
                    <PaginationItem key={p}>
                      <PaginationLink
                        href="#"
                        isActive={p === page}
                        onClick={(e) => {
                          e.preventDefault();
                          handlePageClick(p);
                        }}
                      >
                        {p}
                      </PaginationLink>
                    </PaginationItem>
                  )
                )}

                <PaginationItem>
                  <PaginationNext
                    href="#"
                    onClick={(e) => {
                      e.preventDefault();
                      if (page < totalPages) setPage(page + 1);
                    }}
                  />
                </PaginationItem>
              </PaginationContent>
            </Pagination>
          </div>
        )}
      </main>
    </div>
  );
}
