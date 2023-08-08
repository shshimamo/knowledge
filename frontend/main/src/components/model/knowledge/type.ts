export type Knowledge = {
  id: number;
  userId: number;
  title: string;
  text: string;
  isPublic: boolean;
  publishedAt: Date | null;
}