import { useCallback, useEffect, useRef } from 'react';

export function useSpeech() {
  const voiceRef = useRef<SpeechSynthesisVoice | null>(null);

  useEffect(() => {
    const load = () => {
      const voices = window.speechSynthesis.getVoices();
      voiceRef.current =
        voices.find((v) => v.lang.startsWith('en') && /female|samantha|karen/i.test(v.name)) ??
        voices.find((v) => v.lang.startsWith('en')) ??
        voices[0] ?? null;
    };
    load();
    window.speechSynthesis.onvoiceschanged = load;
    return () => window.speechSynthesis.cancel();
  }, []);

  const speak = useCallback((text: string) => {
    if (!text) return;
    window.speechSynthesis.cancel();
    const u = new SpeechSynthesisUtterance(text);
    if (voiceRef.current) u.voice = voiceRef.current;
    u.rate = 0.9;
    u.pitch = 1.0;
    window.speechSynthesis.speak(u);
  }, []);

  const cancel = useCallback(() => window.speechSynthesis.cancel(), []);

  return { speak, cancel };
}