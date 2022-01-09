from os import name
from fastapi import FastAPI, Path
from typing import Optional
from fastapi.param_functions import Query
from pydantic import BaseModel
from starlette.background import BackgroundTask

app = FastAPI()

class Item(BaseModel):
    name: str
    price: float
    brand: Optional[str] = None

class UpdateItem(BaseModel):
    name: Optional[str] = None
    price: Optional[float] = None
    brand: Optional[str] = None



inventory = {}

@app.get("/get-item/{item_id}")
def get_item(item_id: int = Path(None, description="the ID of the item you want to view.", gt = 0)):
    return inventory[item_id]


@app.get("/get-by-name")
def get_item(name: str = Query(None, title = "Name", description="Name of item.",max_length= 10, min_length=2)):
    for item_id in inventory:
        if inventory[item_id].name == name:
            return inventory[item_id] 
    return {"Data": "Not found"}


@app.post("/create-item/{item_id}")
def create_item(item_id: int, item: Item):
    if item_id in inventory:
        return {"Error": "Item Id already exist"}

    inventory[item_id] = {"name": item.name, "brand": item.brand, "price": item.price}

    return inventory[item_id]


@app.put("/update-item/{item_id}")
def update_item(item_id: int, item: UpdateItem):
    if item_id not in inventory:
        return {"Error": "Item Id doesnt yeahexist"}

    inventory[item_id].name = item.name 
    if item.name != None:
        inventory[item_id].name = item.name

    if item.name != None:
        inventory[item_id].price = item.price

    if item.name != None:
        inventory[item_id].brand = item.brand

    return inventory[item_id]

@app.delete("/delete-item")
def delete_item(item_id: int = Query(..., description="the id of the item to delete",item_idye >= 0))
    return()

