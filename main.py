import telebot
from telebot import types
import threading
import time

age = 0
f1 = False
f2 = False

def wait1(bot):
    global f1
    global f2
    time.sleep(3)
    if f1==False:
       bot.send_message(566423152, 'вы проигнорировали')
       f2 = True


bot = telebot.TeleBot('2010595022:AAG1xMs2CgrcIRSkBJZHvsCgLp3ZnFkkHko')


@bot.message_handler(content_types=['text'])
def start(message):
    if message.text == '/reg':
        bot.send_message(message.from_user.id, "Что-то")
        bot.register_next_step_handler(message, get_age) #следующий шаг – функция get_name
    else:
        bot.send_message(message.from_user.id, 'Напиши /reg')

def get_age(message):
    global f1
    global f2
    global age
    while age == 0: #проверяем что возраст изменился
        try:
             age = int(message.text) #проверяем, что возраст введен корректно
        except Exception:
             bot.send_message(message.from_user.id, 'Цифрами, пожалуйста')

        keyboard = types.InlineKeyboardMarkup() #наша клавиатура
        key_yes = types.InlineKeyboardButton(text='Да', callback_data='yes') #кнопка «Да»
        keyboard.add(key_yes) #добавляем кнопку в клавиатуру
        key_no= types.InlineKeyboardButton(text='Нет', callback_data='no')
        keyboard.add(key_no)
        question = 'Еще что-то'
        bot.send_message(message.from_user.id, text=question, reply_markup=keyboard)
        t1 = threading.Thread(target=wait1, args=(bot,))
        t1.start()
        t1.join()
        return

@bot.callback_query_handler(func=lambda call: True)
def callback_worker(call):
    global f1
    if call.data == "yes": #call.data это callback_data, которую мы указали при объявлении кнопки
        f1 = True #код сохранения данных, или их обработки
        if f2 == True:
            bot.send_message(call.message.chat.id, 'Игнорировать плохо')
        else:
            bot.send_message(call.message.chat.id, 'Запомню : )')
        
    elif call.data == "no":
        f1= True
         #переспрашиваем
    
    return
    
bot.polling(none_stop=True, interval=0)