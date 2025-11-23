import React, { useEffect, useState } from "react";
import { apiGet, apiPost } from "../api";
import { toast } from "react-toastify";

function ItemsPage({ onLogout }) {
  const [items, setItems] = useState([]);
  const [cartId, setCartId] = useState(null);

  const token = localStorage.getItem("token");

  useEffect(() => {
    loadItems();
  }, []);

  async function loadItems() {
    const data = await apiGet("/items");
    setItems(data || []);
  }

  // Add item to cart + toast feedback
  async function addToCart(itemId) {
    const result = await apiPost("/carts", { item_id: itemId }, token);

    if (result && result.cart_id) {
      setCartId(result.cart_id);
      toast.success("Item added to cart");
    } else {
      toast.error("Could not add item to cart");
    }
  }

  // Checkout → convert cart to order + toast
  async function checkout() {
    if (!cartId) {
      toast.warn("Cart is empty");
      return;
    }

    const res = await apiPost("/orders", { cart_id: cartId }, token);

    if (res && res.order_id) {
      toast.success("Order successful");
      setCartId(null);
    } else {
      toast.error("Something went wrong while placing order");
    }
  }

  // Show cart contents (using toast instead of alert)
  async function showCart() {
    const res = await apiGet("/carts", token);

    if (!res || !res.cart_id || !Array.isArray(res.items) || res.items.length === 0) {
      toast.info("Cart is empty");
      return;
    }

    const count = res.items.length;
    const ids = res.items.map((it) => it.ItemID).join(", ");

    toast.info(`Cart ID: ${res.cart_id} | Items (${count}): ${ids}`);
  }

  // Show order history (toast instead of alert)
  async function showOrderHistory() {
    const res = await apiGet("/orders", token);

    if (!Array.isArray(res) || res.length === 0) {
      toast.info("You have no orders yet");
      return;
    }

    const ids = res.map((o) => o.ID).join(", ");
    toast.info(`Your orders: ${ids}`);
  }

  return (
    <div className="item-page">
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <h2>Items</h2>
        <button onClick={onLogout}>Logout</button>
      </div>

      <div className="top-buttons">
        <button onClick={showCart}>Cart</button>
        <button onClick={checkout}>Checkout</button>
        <button onClick={showOrderHistory}>Order History</button>
      </div>

      {items.map((item) => (
        <div key={item.ID} className="item-box">
          <span>{item.Name}</span>
          <button onClick={() => addToCart(item.ID)}>Add</button>
        </div>
      ))}
    </div>
  );
}

export default ItemsPage;
