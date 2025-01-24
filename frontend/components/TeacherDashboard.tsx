"use client";

import { useState, useEffect } from "react";
import { GoogleGenerativeAI } from "@google/generative-ai";
import { useToast } from "@/hooks/use-toast";
import { useApp } from "../contexts/AppContext";
import type { Assignment, Subject } from "../types";
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

export default function TeacherDashboard() {
  const { currentUser, getAssignments, submitGrade } = useApp();
  const [selectedSubject, setSelectedSubject] = useState<Subject | "ALL">(
    "ALL",
  );
  const [assignments, setAssignments] = useState<Assignment[]>([]);
  const [selectedAssignment, setSelectedAssignment] =
    useState<Assignment | null>(null);
  const [feedback, setFeedback] = useState("");
  const [grade, setGrade] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isSubmittingGrade, setIsSubmittingGrade] = useState(false);
  const [isGeneratingFeedback, setIsGeneratingFeedback] = useState(false);

  const { toast } = useToast();

  const generateAIFeedback = async (assignment: Assignment) => {
    if (!process.env.NEXT_PUBLIC_GEMINI_API_KEY) {
      toast({
        variant: "destructive",
        title: "API Key Missing",
        description:
          "Gemini API key is not configured. AI feedback generation is unavailable.",
      });
      return;
    }

    setIsGeneratingFeedback(true);
    try {
      const genAI = new GoogleGenerativeAI(
        process.env.NEXT_PUBLIC_GEMINI_API_KEY,
      );
      const model = genAI.getGenerativeModel({ model: "gemini-pro" });

      const prompt = `You are an experienced teacher assisting in grading assignments.
      Please provide constructive feedback for the following ${assignment.subject} assignment:

      Title: ${assignment.title}
      Content: ${assignment.content}

      Please provide:
      1. A brief overall assessment
      2. 2-3 specific strengths
      3. 2-3 areas for improvement
      4. Suggested grade (0-100)

      Keep the tone encouraging and constructive.`;

      const result = await model.generateContent(prompt);
      const response = result.response;
      const generatedFeedback = response.text();

      toast({
        title: "Feedback Generated",
        description: "AI feedback has been generated successfully.",
      });
      setFeedback(generatedFeedback);
      // Extract suggested grade from the feedback if available
      const gradeMatch = generatedFeedback.match(/\b([0-9]{1,3})\b/);
      if (gradeMatch) {
        setGrade(gradeMatch[0]);
      }
    } catch (error) {
      console.error("Failed to generate AI feedback:", error);
      toast({
        variant: "destructive",
        title: "Generation Failed",
        description: "Failed to generate AI feedback. Please try again later.",
      });
    } finally {
      setIsGeneratingFeedback(false);
    }
  };

  useEffect(() => {
    loadAssignments(selectedSubject);
  }, [selectedSubject]);

  const loadAssignments = async (subject?: Subject | "ALL") => {
    setIsLoading(true);
    try {
      subject = subject === "ALL" ? undefined : subject;
      const fetchedAssignments = await getAssignments(subject);
      setAssignments(fetchedAssignments);
    } catch (error) {
      console.error("Failed to load assignments:", error);
      // alert("Failed to load assignments. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  const handleGrade = async () => {
    if (selectedAssignment) {
      setIsSubmittingGrade(true);
      try {
        await submitGrade(selectedAssignment.id, Number(grade), feedback);
        // alert("Grade submitted successfully!");
        setSelectedAssignment(null);
        setFeedback("");
        setGrade("");
        await loadAssignments(selectedSubject);
      } catch (error) {
        console.error("Failed to submit grade:", error);
        // alert("Failed to submit grade. Please try again.")
      } finally {
        setIsSubmittingGrade(false);
      }
    }
  };

  if (currentUser?.role !== "TEACHER") return null;

  return (
    <div className="mt-4">
      <h2 className="text-xl font-bold mb-2">Teacher Dashboard</h2>
      <Select
        value={selectedSubject}
        onValueChange={(value: Subject) => setSelectedSubject(value)}
      >
        <SelectTrigger className="w-[180px]">
          <SelectValue placeholder="Select subject" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="ALL">All Subjects</SelectItem>
          <SelectItem value="ENGLISH">English</SelectItem>
          <SelectItem value="MATH">Math</SelectItem>
        </SelectContent>
      </Select>
      <div className="mt-4 grid grid-cols-2 gap-4">
        <div>
          <h3 className="text-lg font-semibold mb-2">Assignments</h3>
          {isLoading ? (
            <p>Loading assignments...</p>
          ) : !assignments || assignments.length === 0 ? (
            <p className="text-gray-500 italic">
              No assignments available for this subject
            </p>
          ) : (
            assignments.map((assignment) => (
              <div key={assignment.id} className="mb-2 p-2 border rounded">
                <p>
                  <strong>{assignment.title}</strong> by Student ID:{" "}
                  {assignment.student_id}
                </p>
                <Button onClick={() => setSelectedAssignment(assignment)}>
                  Review
                </Button>
              </div>
            ))
          )}
        </div>
        {selectedAssignment && (
          <div>
            <h3 className="text-lg font-semibold mb-2">Grade Assignment</h3>
            <p>
              <strong>Title:</strong> {selectedAssignment.title}
            </p>
            <p>
              <strong>Student ID:</strong> {selectedAssignment.student_id}
            </p>
            <p>
              <strong>Content:</strong> {selectedAssignment.content}
            </p>
            <Button
              className="mt-2 mb-4"
              onClick={() => generateAIFeedback(selectedAssignment)}
              disabled={isGeneratingFeedback}
            >
              {isGeneratingFeedback
                ? "Generating Feedback..."
                : "Generate AI Feedback"}
            </Button>
            <Textarea
              className="mt-2"
              placeholder="Feedback"
              value={feedback}
              onChange={(e) => setFeedback(e.target.value)}
            />
            <Input
              className="mt-2"
              type="number"
              placeholder="Grade"
              value={grade}
              onChange={(e) => setGrade(e.target.value)}
            />
            <Button
              className="mt-2"
              onClick={handleGrade}
              disabled={isSubmittingGrade}
            >
              {isSubmittingGrade ? "Submitting Grade..." : "Submit Grade"}
            </Button>
          </div>
        )}
      </div>
    </div>
  );
}
