"use client";

import { useState, useEffect } from "react";
import { useApp } from "../contexts/AppContext";
import { type Subject, type Grade } from "../types";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export default function StudentDashboard() {
  const { currentUser, submitAssignment, getGrades } = useApp();
  const [subject, setSubject] = useState<Subject>("ENGLISH");
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [grades, setGrades] = useState<Grade[]>([]);
  const [isLoadingGrades, setIsLoadingGrades] = useState(false);

  useEffect(() => {
    if (currentUser) {
      loadGrades();
    }
  }, [currentUser]);

  const loadGrades = async () => {
    if (currentUser) {
      setIsLoadingGrades(true);
      try {
        const fetchedGrades = await getGrades(currentUser.id);
        setGrades(fetchedGrades);
      } catch (error) {
        console.error("Failed to load grades:", error);
        // alert("Failed to load grades. Please try again.")
      } finally {
        setIsLoadingGrades(false);
      }
    }
  };

  const handleSubmit = async () => {
    setIsSubmitting(true);
    try {
      await submitAssignment(title, subject, content);
      setTitle("");
      setContent("");
      // Fetch grades again to update the list
      await loadGrades();
      // alert("Assignment submitted successfully!");
    } catch (error) {
      console.error("Failed to submit assignment:", error);
      // alert("Failed to submit assignment. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  };

  if (!currentUser || currentUser.role !== "STUDENT") return null;

  return (
    <div className="mt-4">
      <h2 className="text-xl font-bold mb-2">Student Dashboard</h2>
      <div className="grid grid-cols-2 gap-4">
        <div>
          <h3 className="text-lg font-semibold mb-2">Submit Assignment</h3>
          <Select
            value={subject}
            onValueChange={(value: Subject) => setSubject(value)}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select subject" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="ENGLISH">English</SelectItem>
              <SelectItem value="MATH">Math</SelectItem>
            </SelectContent>
          </Select>
          <Input
            className="mt-2"
            placeholder="Assignment Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <Textarea
            className="mt-2"
            placeholder="Assignment Content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
          <Button
            className="mt-2"
            onClick={handleSubmit}
            disabled={isSubmitting}
          >
            {isSubmitting ? "Submitting..." : "Submit Assignment"}
          </Button>
        </div>
        <div>
          <h3 className="text-lg font-semibold mb-2">My Grades</h3>
          {isLoadingGrades ? (
            <p>Loading grades...</p>
          ) : !grades || grades.length === 0 ? (
            <p className="text-gray-500 italic">No grades available yet</p>
          ) : (
            grades.map((grade) => (
              <div key={grade.id} className="mb-2 p-2 border rounded">
                <p>
                  <strong>Assignment ID:</strong> {grade.assignment_id}
                </p>
                <p>
                  <strong>Grade:</strong> {grade.grade}
                </p>
                <p>
                  <strong>Feedback:</strong> {grade.feedback}
                </p>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
