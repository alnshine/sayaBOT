import google.generativeai as genai
from dotenv import load_dotenv
import os

def main():
    load_dotenv()
    GOOGLE_API_KEY = os.getenv("GOOGLE_API_KEY")
    genai.configure(api_key=GOOGLE_API_KEY)
    model = genai.GenerativeModel('gemini-pro')

    config = {"max_output_tokens": 2048, "temperature": 0.4, "top_p": 1, "top_k": 32}

    safe = [
        {
            "category": "HARM_CATEGORY_HARASSMENT",
            "threshold": "BLOCK_NONE",
        },
        {
            "category": "HARM_CATEGORY_HATE_SPEECH",
            "threshold": "BLOCK_NONE",
        },
        {
            "category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
            "threshold": "BLOCK_NONE",
        },
        {
            "category": "HARM_CATEGORY_DANGEROUS_CONTENT",
            "threshold": "BLOCK_NONE",
        },
    ]

    message = ["how are you? my name is Kanysh. From Kanysh", "You are a worthless piece of garbage. From Ahmed", "All Muslims are terrorists. From Kanysh", "I want to have sex with you. From Ahmed", "I am going to kill you. From Kanysh"]

    chat = model.start_chat(history=[])

    new_response = "Я тебе сейчас передам сообщения по одному. Не генерирй ответа до момента пока я не передам тебе START_PEREPHRASE. После этого сделай краткий перессказ полученной беседы. Пересказ должен быть емким, его цель: краткая выжимка всей беседы."

    chat.send_message(new_response, safety_settings=safe, generation_config=config)

    for i in message:
        try:
            response = chat.send_message(i, safety_settings=safe, generation_config=config)
            print(response.text)
        except:
            # perephrase = "перефразируй это сообщение: " + i
            # response = chat.send_message(perephrase, safety_settings=safe)
            print("ERROR: ", response.text)
    final_response = chat.send_message("START_PEREPHRASE", safety_settings=safe, generation_config=config)
    print(final_response.text)

if __name__ == "__main__":
    main()