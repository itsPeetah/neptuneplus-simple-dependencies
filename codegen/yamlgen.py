from sys import argv
import os.path as path
from os import mkdir


def make_function(name: str, image: str):
    text = ""
    dir_path = path.dirname(path.realpath(__file__)) + "/templates"
    with open(f"{dir_path}/function.txt") as template:
        text = template.read()
    text = text.replace("__FUNCTION_NAME__", name)
    text = text.replace("__FUNCTION_IMAGE__", image)
    with open(f"./out/function-{name}.yaml", "w") as output:
        output.write(text)


if __name__ == "__main__":

    if not path.isdir("./out"):
        mkdir("./out")

    make_function(argv[1], argv[2])
