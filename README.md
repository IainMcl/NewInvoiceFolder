# NewInvoiceFolder
Simple exe to create a new invoice folder

> Why spend 10s doing a task that you can spend hours failing to automate

## Issue

I currently create invoices which are in a folder with `Invoice 1`, `Invoice 2`, ... These all come from a common template folder `Invoice_template` containing an `xlsx` file. This structure is used by another script to read invoices and pull together my tax calculations. 

To create a new invoice I have to copy the `Invoice_template` folder and rename the coppied folder to the latest invoice number. I also have to rename the template `xlsx` file that it contains.

## Solution

Rather than doing this 10s task over and over again I now have an `exe` file that I can double click in my invoice folder which will copy the templates and rename them appropreatly.

# How to run

In the folder run 

```bash
go build .
```

This creates a .exe file which can be coppied into the invoice folder.

Double click the .exe file in the invoice folder and a new invoice folder will be coppied from the template.

# Note

Folders and template folders only exist for testing
