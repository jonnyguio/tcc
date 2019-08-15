library(ggplot2)
library(plyr)
library(dplyr)
library(gmodels)

temp = list.files(path = "./go-with-branching/", pattern = "*.csv.fix", full.names = TRUE)
go_with_branching_dataset = lapply(temp, read.delim, sep = ",")
go_with_branching_merged_dataset = rbind.fill(go_with_branching_dataset)

temp = list.files(path = "./c-with-branching/", pattern = "*.csv", full.names = TRUE)
c_with_branching_dataset = lapply(temp, read.delim, sep = ",")
c_with_branching_merged_dataset = rbind.fill(c_with_branching_dataset)

temp = list.files(path = "./go-cpp-algorithm/", pattern = "*.csv.fix", full.names = TRUE)
go_cpp_branching_dataset = lapply(temp, read.delim, sep = ",")
go_cpp_branching_merged_dataset = rbind.fill(go_cpp_branching_dataset)

colours_cores <- c("#990000", "#274e13", "#073763")
substr(colours_cores, 1, 7)
colours_cores[1]
standard_error <- function(x) sd(x)/sqrt(length(x))
if (FALSE) {
  go_no_branching_summary <- go_no_branching_merged_dataset %>%
    group_by(cores, matrix_size) %>%
    summarise(mean_time=mean(time), sd_time=sd(time), cv=(sd(time) / mean(time))*100, std_err=standard_error(time), min_time=min(time), max_time=max(time))
}
go_with_branching_summary <- go_with_branching_merged_dataset %>%
  group_by(cores, matrix_size) %>%
  summarise(mean_time=mean(time), sd_time=sd(time), cv=(sd(time) / mean(time))*100, std_err=standard_error(time), min_time=min(time), max_time=max(time))
if (FALSE) {
  c_no_branching_summary <- c_no_branching_merged_dataset %>%
    group_by(cores, matrix_size) %>%
    summarise(mean_time=mean(time), sd_time=sd(time), cv=(sd(time) / mean(time))*100, std_err=standard_error(time), min_time=min(time), max_time=max(time))
}
c_with_branching_summary <- c_with_branching_merged_dataset %>%
  group_by(cores, matrix_size) %>%
  summarise(mean_time=mean(time), sd_time=sd(time), cv=(sd(time) / mean(time))*100, std_err=standard_error(time), min_time=min(time), max_time=max(time))

go_cpp_branching_summary <- go_cpp_branching_merged_dataset %>%
  group_by(cores, matrix_size) %>%
  summarise(mean_time=mean(time), sd_time=sd(time), cv=(sd(time) / mean(time))*100, std_err=standard_error(time), min_time=min(time), max_time=max(time))

# c_go_no_branching_comparison = rbind(c_no_branching_summary, go_no_branching_summary)
# c_go_no_branching_comparison$lang = ""
# c_go_no_branching_comparison$lang[1:18] = "C++"
# c_go_no_branching_comparison$lang[19:36] = "Go"
# c_go_no_branching_comparison = c_go_no_branching_comparison[c_go_no_branching_comparison$cores == 4,]


c_go_with_branching_comparison = rbind(c_with_branching_summary, go_with_branching_summary)
c_go_with_branching_comparison$lang = ""
c_go_with_branching_comparison$lang[1:18] = "C++"
c_go_with_branching_comparison$lang[19:36] = "Go"
c_go_with_branching_comparison = c_go_with_branching_comparison[c_go_with_branching_comparison$cores == 4,]

c_go_same_algorithm_comparison = rbind(c_with_branching_summary, go_cpp_branching_summary)
c_go_same_algorithm_comparison$lang = ""
c_go_same_algorithm_comparison$lang[1:18] = "C++"
c_go_same_algorithm_comparison$lang[19:36] = "Go"
c_go_same_algorithm_comparison = c_go_same_algorithm_comparison[c_go_same_algorithm_comparison$cores == 4,]


if (FALSE) {
ggplot(mapping = aes(x = go_no_branching_summary$matrix_size,
                     y = go_no_branching_summary$mean_time,
                     colour = factor(go_no_branching_summary$cores))) +
       geom_point() +
       geom_line() + 
       geom_errorbar(aes(ymin=go_no_branching_summary$min_time, ymax=go_no_branching_summary$max_time), width = 100) +
       scale_color_manual(values = colours_cores) +
       labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Núcleos")
}

current_dataset <- go_with_branching_summary
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = factor(current_dataset$cores))) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Núcleos")

if (FALSE) {
current_dataset <- c_no_branching_summary
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = factor(current_dataset$cores))) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Núcleos")
}

current_dataset <- c_with_branching_summary
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = factor(current_dataset$cores))) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Núcleos")

current_dataset <- go_cpp_branching_summary
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = factor(current_dataset$cores))) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Núcleos")

# Comparacões
if (FALSE) {
current_dataset <- c_go_no_branching_comparison
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = current_dataset$lang)) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Linguagem")
}

current_dataset <- c_go_with_branching_comparison
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = current_dataset$lang)) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Linguagem")


current_dataset <- c_go_same_algorithm_comparison
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = current_dataset$lang)) +
  geom_point() +
  geom_line() + 
  geom_errorbar(aes(ymin=current_dataset$min_time, ymax=current_dataset$max_time), width = 100) +
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Linguagem")


# Comparacões sem error bar
if (FALSE) {
current_dataset <- c_go_no_branching_comparison
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = current_dataset$lang)) +
  geom_point() +
  geom_line() + 
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Linguagem", title = "Comparacão Go vs C++ sem otimização - 4 núcleos")
}

current_dataset <- c_go_with_branching_comparison
ggplot(mapping = aes(x = current_dataset$matrix_size,
                     y = current_dataset$mean_time,
                     colour = current_dataset$lang)) +
  geom_point() +
  geom_line() + 
  scale_color_manual(values = colours_cores) +
  labs(x = "Tamanho da Matriz", y = "Tempo (s)", colour = "Linguagem", title = "Comparacão Go vs C++ com otimização - 4 núcleos")


################## ANTIGO
#################

geom_errorbar()
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
                  # colour = factor(current_dataset[[current_row]]$cores)),
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
