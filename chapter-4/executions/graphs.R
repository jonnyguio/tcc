library(ggplot2)

colours_cores <- c("#990000", "#274e13", "#073763")

substr(colours_cores, 1, 7)
colours_cores[1]

ggplot(mapping = aes()) +
  geom_smooth(aes(x = c_with_branching_dataset_1$sample_number,
                  y = c_with_branching_dataset_1$time,
                  colour = factor(c_with_branching_dataset_1$cores)),
              c_with_branching_dataset_1,
              fill = colours_cores[1]) + 
  geom_smooth(aes(x = c_with_branching_dataset_2$sample_number,
                  y = c_with_branching_dataset_2$time,
                  colour = factor(c_with_branching_dataset_2$cores)),
              c_with_branching_dataset_2,
              fill = colours_cores[2]) + 
  geom_smooth(aes(x = c_with_branching_dataset_4$sample_number,
                  y = c_with_branching_dataset_4$time,
                  colour = factor(c_with_branching_dataset_4$cores)),
              c_with_branching_dataset_4,
              fill = colours_cores[3]) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "C - Matriz 5000x5000")

current_dataset = c_no_branching_dataset
current_row <- 6
ggplot(mapping = aes()) +
  geom_smooth(aes(x = current_dataset[[current_row]]$sample_number,
                  y = current_dataset[[current_row]]$time,
                  colour = factor(current_dataset[[current_row]]$cores)),
              current_dataset[[current_row]],
              fill = colours_cores[1]) + 
  geom_smooth(aes(x = current_dataset[[current_row + 6]]$sample_number,
                  y = current_dataset[[current_row + 6]]$time,
                  colour = factor(current_dataset[[current_row + 6]]$cores)),
              current_dataset[[current_row + 6]],
              fill = colours_cores[2]) + 
  geom_smooth(aes(x = current_dataset[[current_row + 12]]$sample_number,
                  y = current_dataset[[current_row + 12]]$time,
                  colour = factor(current_dataset[[current_row + 12]]$cores)),
              current_dataset[[current_row + 12]],
              fill = colours_cores[3]) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "C - Matriz 5000x5000")

current_dataset = c_with_branching_datasets
current_row <- 1
ggplot(mapping = aes()) +
  geom_smooth(aes(x = current_dataset[[current_row]]$sample_number,
                  y = current_dataset[[current_row]]$time,
                  colour = factor(current_dataset[[current_row]]$cores)),
              current_dataset[[current_row]],
              fill = colours_cores[1], se = FALSE) + 
  geom_smooth(aes(x = current_dataset[[current_row + 6]]$sample_number,
                  y = current_dataset[[current_row + 6]]$time,
                  colour = factor(current_dataset[[current_row + 6]]$cores)),
              current_dataset[[current_row + 6]],
              fill = colours_cores[2], se = FALSE) + 
  geom_smooth(aes(x = current_dataset[[current_row + 12]]$sample_number,
                  y = current_dataset[[current_row + 12]]$time,
                  colour = factor(current_dataset[[current_row + 12]]$cores)),
              current_dataset[[current_row + 12]],
              fill = colours_cores[3], se = FALSE) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "C com branching - Matriz 100x100")

current_dataset = c_with_branching_datasets
current_row <- 6
ggplot(mapping = aes()) +
  geom_smooth(aes(x = current_dataset[[current_row]]$sample_number,
                  y = current_dataset[[current_row]]$time,
                  colour = factor(current_dataset[[current_row]]$cores)),
              current_dataset[[current_row]],
              fill = colours_cores[1]) + 
  geom_smooth(aes(x = current_dataset[[current_row + 6]]$sample_number,
                  y = current_dataset[[current_row + 6]]$time,
                  colour = factor(current_dataset[[current_row + 6]]$cores)),
              current_dataset[[current_row + 6]],
              fill = colours_cores[2]) + 
  geom_smooth(aes(x = current_dataset[[current_row + 12]]$sample_number,
                  y = current_dataset[[current_row + 12]]$time,
                  colour = factor(current_dataset[[current_row + 12]]$cores)),
              current_dataset[[current_row + 12]],
              fill = colours_cores[3]) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "C com branching - Matriz 4000x4000")

current_dataset = go_no_branching_dataset
current_row <- 1
ggplot(mapping = aes()) +
  geom_boxplot(aes(x = current_dataset[[current_row + 6]]$sample_number,
                   y = current_dataset[[current_row + 6]]$time,
                   colour = factor(current_dataset[[current_row]]$cores)),
               current_dataset[[current_row]]) +
  geom_boxplot(aes(x = current_dataset[[current_row + 6]]$sample_number,
                  y = current_dataset[[current_row + 6]]$time,
                  colour = factor(current_dataset[[current_row + 6]]$cores)),
              current_dataset[[current_row + 6]]) + 
  geom_boxplot(aes(x = current_dataset[[current_row + 12]]$sample_number,
                  y = current_dataset[[current_row + 12]]$time,
                  colour = factor(current_dataset[[current_row + 12]]$cores)),
              current_dataset[[current_row + 12]]) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "C com branching - Matriz 4000x4000")




ggplot(mapping = aes()) +
  geom_smooth(aes(x = execution.1.1000.csv$sample_number,
                  y = execution.1.1000.csv$time,
                  colour = factor(execution.1.1000.csv$cores)),
              execution.1.1000.csv,
              fill = colours_cores[1]) + 
  geom_smooth(aes(x = execution.2.1000.csv$sample_number,
                  y = execution.2.1000.csv$time,
                  colour = factor(execution.2.1000.csv$cores)),
              execution.2.1000.csv,
              fill = colours_cores[2]) + 
  geom_smooth(aes(x = execution.4.1000.csv$sample_number,
                  y = execution.4.1000.csv$time,
                  colour = factor(execution.4.1000.csv$cores)),
              execution.4.1000.csv,
              fill = colours_cores[3]) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "Matriz 1000x1000")



ggplot(mapping = aes()) +
  geom_smooth(aes(x = execution.1.5000.csv$sample_number,
                  y = execution.1.5000.csv$time,
                  colour = factor(execution.1.5000.csv$cores)),
              execution.1.100.csv,
              fill = colours_cores[1]) + 
  geom_smooth(aes(x = execution.2.5000.csv$sample_number,
                  y = execution.2.5000.csv$time,
                  colour = factor(execution.2.5000.csv$cores)),
              execution.2.5000.csv,
              fill = colours_cores[2]) + 
  geom_smooth(aes(x = execution.4.5000.csv$sample_number,
                  y = execution.4.5000.csv$time,
                  colour = factor(execution.4.5000.csv$cores)),
              execution.4.5000.csv,
              fill = colours_cores[3]) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Amostra", y = "Tempo (s)", colour = "Núcleos", title = "Matriz 5000x5000")
