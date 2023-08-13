declare global {
  export interface Window {
    Go: any;
    drawImages: (img1: string, img2: string) => string
  }
}

export {};
