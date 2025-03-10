import tkinter as tk
from tkinter import ttk

def display_table(val1, val2, val3):
    # Create the main window with a larger size
    root = tk.Tk()
    root.title("3x2 Table")
    root.geometry("500x300")  # Set a larger window size
    
    # Define a larger font for the labels
    label_font = ("Helvetica", 16)
    
    # Define the fixed labels for the first column
    row_labels = ["CL", "cl", "CLI"]
    # Define the values for the second column from the function arguments
    values = [val1, val2, val3]
    
    # Create a 3x2 table using grid layout
    for i in range(3):
        # First column: fixed label
        label = ttk.Label(root, text=row_labels[i], borderwidth=1, relief="solid", padding=10, font=label_font)
        label.grid(row=i, column=0, padx=5, pady=5, sticky="nsew")
        
        # Second column: value passed as argument
        value_label = ttk.Label(root, text=str(values[i]), borderwidth=1, relief="solid", padding=10, font=label_font)
        value_label.grid(row=i, column=1, padx=5, pady=5, sticky="nsew")
    
    # Make columns expand equally
    root.grid_columnconfigure(0, weight=1)
    root.grid_columnconfigure(1, weight=1)
    
    # Run the GUI event loop
    root.mainloop()
