# CSU-Industry-Skills

Research Project for Comparing CSU and Industry Skills

## Running BERT Model

`gaps_bert.ipynb` can be opened using VS Code or Jupyter Notebook software.

It is recommended to create a conda environment for the notebook, using Python 3.11

### Data Processing

If you have the file `.\bert-model\data\bert_training_data.csv`, skip this section.

If you don't have the file `.\bert-model\data\bert_training_data.csv`, uncomment the code in cells 2, 3, 8, and 9 to create the file.

- Be sure to have `descriptions.txt` and `all-csu-codes.csv` within the folder `bert-model/data/` when executing a fresh training session.
- Once the training data is created, comment out those cells again.

### Package Installationspip install pandas

Install necessary packages: 

`pip`: 

- `transformers`
- `wandb`
- `pandas`
- `datasets`
- `sentence-transformers`

`conda`:

- `scikit-learn`
- Install Pytorch versioning that matches your device

### Weights and Biases Set-up

Create or log into an account with [Weights & Biases](wandb.ai), to retrieve an API key.
 - Place API key in a txt file named `WANDB_API_KEY.txt`, located at `./dpr-model/WANDB_API_KEY.txt`.
 - Change the `entity` parameter of wandb.init(), within the training function to match your wandb account entity name.

## Directory Descriptions
Due to the limitation of Github, we were unable to upload our Atlas TI projects into this repository. Some of the code in this repo are associated with converting text files to CSVs for `DrawIO`.  
### CSU-CODES
This directory contains both a `csv` and `xlsx` prepresentation of each course number, and all codes associated to said course. Here is an example: 
```
CS455,Abstraction,Algorithms Underpinning P2P Sys...
```
### DRAWIO-TEXT
This was the landing space for converting the Atlat TI exported data into readable text files to be sent to DrawIO to create visual graph representation. 
### INDUSTRY-CODES
This is similar to `CSU-CODES` but it is associated with the industry jobs instead of CSU CS corses.
### LIB
This is the brains of the conversion. `exel_to_csv.go` will convert the xlsx files exported from Atlas TI and will convert them into CSVs (better for applying pandas to them). Also `draw_io_converter.ipynd` will convert each CSV file to a readable text file that drawio's API can convert from a text to a visual graph.
Secondly, `go.mod` and `go.sum` were used in the go file `exel_to_csv.go`.

