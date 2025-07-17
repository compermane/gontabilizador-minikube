#!/usr/bin/env python3

import mysql.connector
import time
from datetime import datetime

config = {
    'user': 'eugenio',
    'password': 'aA@12345',
    'host': 'mysql',
    'database': 'gontabilizador'
}

def get_latest_update_time(cursor, cnx):
    cnx.commit()
    cursor.execute("SELECT MAX(updated_at) FROM presenca")
    result = cursor.fetchone()
    return result[0] if result[0] else datetime.min

def main():
    try:
        cnx = mysql.connector.connect(**config)
        cursor = cnx.cursor()
        print("Conectado ao banco MySQL")
    except Exception as e:
        print(f"Erro ao conectar ao MySQL: {e}")
        return

    last_checked = datetime.min

    while True:
        try:
            cursor = cnx.cursor()
            latest_update = get_latest_update_time(cursor, cnx)
            cursor.close()
            if latest_update > last_checked:
                print(f"Novas atualizações detectadas em presenca em {latest_update}")
                last_checked = latest_update
            else:
                print("Sem alterações recentes.")
        except Exception as e:
            print(f"Erro ao consultar banco: {e}")

        time.sleep(5)

if __name__ == "__main__":
    print("Iniciando polling...")
    main()