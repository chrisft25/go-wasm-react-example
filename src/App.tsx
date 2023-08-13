import { Input } from "./@/src/components/ui/input"
import { Label } from "./@/src/components/ui/label"
import React, { useState, useEffect } from "react"

function App() {

  const [images, setImages] = useState({ img1: null, img2: null })
  const [result, setResult] = useState("")

  async function handleChange(e: React.ChangeEvent<HTMLInputElement>){

    if(!e.target.files?.length) return

    const file = e.target.files[0]

    const data = await toBase64(file)

    setImages({
      ...images,
      [e.target.id]: data
    })
  }

  const toBase64 = (file: File) => new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = reject;
});

  useEffect(()=>{
    if(!images.img1 || !images.img2) return

    const mergedImage = window.drawImages(images.img1, images.img2)
    setResult(mergedImage)
  },[images])

  return (
    <div className='App'>
      <div className="grid w-full max-w-sm items-center gap-1.5">
        <Label htmlFor="img1">Picture 1</Label>
        <Input id="img1" type="file" onChange={handleChange} accept=".png, .jpg, .jpeg"/>
      </div>
      <div className="grid w-full max-w-sm items-center gap-1.5">
        <Label htmlFor="img1">Picture 2</Label>
        <Input id="img2" type="file" onChange={handleChange} accept=".png, .jpg, .jpeg"/>
      </div>
      {
        result.length > 1 && <img src={result} alt="resultado"></img>
      }
    </div>
  );
}

export default App;
