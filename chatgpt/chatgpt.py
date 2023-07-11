import openai
import requests



def github_acquisition():
    use_lang = ""
    language_ratios = {}
    #urlからgithubの情報を抜き出す処理
    url = 'http://localhost:8081/artivles'
    response = requests.get(url)
    repo_data = response.json()
    lang_array = repo_data["lang"]

    for item in lang_array:
        language = item['language']
        ratio = item['ratio']
        language_ratios[language] = ratio
    print(language_ratios)

    for language, ratio in language_ratios.items():
        print(f"{language}: {ratio}%")
        use_lang +=" " + str(language) + ": " + str(ratio) + "%"

    return use_lang




def openai_function(company,field, hobby):

    lang = github_acquisition()
    openai.api_key = ""

    response = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": "あなたは自己PRと志望動機を生成するマシンです。"},
            {"role": "user", "content": "以下は私が志望する会社と分野そして趣味です" + company + "," + field + "," + hobby},
            # {"role": "user", "content": company + "," + field + "," + hobby},
            {"role": "user", "content":  "以下は私のgithubの言語使用割合です。" + lang},
            {"role": "user", "content": "以上の情報で300文字から400字以内で自己PR文と志望動機を作成してください"},
        #     {"role": "user", "content": "項目ごとに何を書いているか記載してください。　例　　以下は自己PR文です"},
         ]
    )

    # print(f"ChatGPT:{response['choices'][0]['message']['content']}")
    return response['choices'][0]['message']['content']


if __name__ == '__main__':
   openai_function("yahoo","インフラ", "IoT")